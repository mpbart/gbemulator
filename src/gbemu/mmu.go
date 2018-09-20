package main

import (
	"fmt"
)

// I/O Registers:
// 0xFF00 - P1   - Joypad input
// 0xFF01 - SB   - Serial transfer data
// 0xFF02 - SC   - Serial I/O control
// 0xFF04 - DIV  - Divider register
// 0xFF05 - TIMA - Timer counter
// 0xFF06 - TMA  - Timer modulo
// 0xFF07 - TAC  - Timer control
// 0xFF0F - IF   - Interrupt flags
// 0xFF10 - NR10 - Sound mode 1 sweep register
// 0xFF11 - NR11 - Sound mode 1 length/wave pattern duty
// 0xFF12 - NR12 - Sound mode 1 envelope
// 0xFF13 - NR13 - Sound mode 1 frequency low
// 0xFF14 - NR14 - Sound mode 1 frequency high
// 0xFF16 - NR21 - Sound mode 2 length/wave pattern duty
// 0xFF17 - NR22 - Sound mode 2 envelope
// 0xFF18 - NR23 - Sound mode 2 frequency low
// 0xFF19 - NR24 - Sound mode 2 frequency high
// 0xFF1A - NR30 - Sound mode 3 on/off
// 0xFF1B - NR31 - Sound mode 3 length/wave pattern duty
// 0xFF1C - NR32 - Sound mode 3 output select
// 0xFF1D - NR33 - Sound mode 3 frequency low
// 0xFF1E - NR34 - Sound mode 3 frequency high
// 0xFF20 - NR41 - Sound mode 4 length/wave pattern duty
// 0xFF21 - NR42 - Sound mode 4 envelope
// 0xFF22 - NR43 - Sound mode 4 polynomial counter
// 0xFF23 - NR44 - Sound mode 4 counter
// 0xFF24 - NR50 - Channel control
// 0xFF25 - NR51 - Sound output selection
// 0xFF26 - NR52 - Sound on/off
// 0xFF30-0xFF3F - Wave pattern RAM
// 0xFF40 - LCDC - LCD Control
// 0xFF41 - STAT - LCDC Status
// 0xFF42 - SCY  - Scroll Y
// 0xFF43 - SCX  - Scroll X
// 0xFF44 - LY   - LCDC Y-Coordinate
// 0xFF45 - LYC  - LY Compare
// 0xFF46 - DMA  - DMA transfer and address
// 0xFF47 - BGP  - BG & Window pallette data
// 0xFF48 - OBP0 - Object pallette 0
// 0xFF49 - OBP1 - Object pallette 1
// 0xFF4A - WY   - Window Y Position
// 0xFF4B - WX   - Window X Position
// 0xFFFF - IE   - Interrupt Enable

type Interrupt int

const (
	VBLANK_INTERRUPT Interrupt = iota
	LCDC_STATUS_INTERRUPT
	TIMER_INTERRUPT
	SERIAL_INTERRUPT
	JOYPAD_INTERRUPT
)

type MMU interface {
	Reset()
	ReadAt(uint16) uint8
	WriteByte(uint16, uint8)
	LCDStatusMode() uint8
	SetLCDStatusMode(uint8)
	SpriteSize() int
	LCDEnabled() bool
	WindowTileMap() uint16
	WindowDisplayEnabled() bool
	BGAndWindowAddressMode() AddressMode
	BGTileMap() uint16
	SpritesEnabled() bool
	BGDisplayPriority() bool
	ConvertNumToPixel(int) RGBPixel
	DisableInterrupts()
	HasPendingInterrupt() bool
	GetNextPendingInterrupt() uint16
	ClearHighestInterrupt()
	FireInterrupt(Interrupt)
	Tick()
}

type mmu struct {
	ROM              [32768]uint8 // 0x0000 - 0x7FFF
	VRAM             [8192]uint8  // 0x8000 - 0x9FFF
	SwitchableRAM    [8192]uint8  // 0xA000 - 0xBFFF
	InternalRAM      [8192]uint8  // 0xC000 - 0xDFFF
	EchoRAM          [8192]uint8  // 0xE000 - 0xFDFF
	OAM              [160]uint8   // 0xFE00 - 0xFE9F
	Unused           [95]uint8    // 0xFEA0 - 0xFEFF
	IoPorts          [128]uint8   // 0xFF00 - 0xFF7F
	HRAM             [127]uint8   // 0xFF80 - 0xFFFE
	InterruptEnable  uint8        // 0xFFFF
	colorMapping     map[int]RGBPixel
	interruptMapping map[int]uint16
}

func CreateMMU() MMU {
	return &mmu{
		colorMapping:     createColorMapping(),
		interruptMapping: createBitToInterruptMap(),
	}
}

func createColorMapping() map[int]RGBPixel {
	m := make(map[int]RGBPixel)
	m[0] = WHITE()
	m[1] = LIGHT_GRAY()
	m[2] = DARK_GRAY()
	m[3] = BLACK()
	return m
}

func createBitToInterruptMap() map[int]uint16 {
	m := make(map[int]uint16)
	m[0] = 0x40
	m[1] = 0x48
	m[2] = 0x50
	m[3] = 0x58
	m[4] = 0x60
	return m
}

func (m *mmu) Reset() {
}

func (m *mmu) ReadAt(address uint16) uint8 {
	switch {
	case address >= 0x0000 && address <= 0x7FFF:
		return m.ROM[address]
	case address >= 0x8000 && address <= 0x9FFF:
		/*
			if !m.CanAccessVRAM() {
				fmt.Println("Accessing VRAM at an illegal time...")
				return 0xFF
			}
		*/
		return m.VRAM[address-0x8000]
	case address >= 0xA000 && address <= 0xBFFF:
		return m.SwitchableRAM[address-0xA000]
	case address >= 0xC000 && address <= 0xDFFF:
		return m.InternalRAM[address-0xC000]
	case address >= 0xE000 && address <= 0xFDFF:
		return m.EchoRAM[address-0xE000]
	case address >= 0xFE00 && address <= 0xFE9F:
		/*
			if !m.CanAccessOAM() {
				fmt.Println("Accessing OAM at an illegal time...")
				return 0xFF
			}
		*/
		return m.OAM[address-0xFE00]
	case address >= 0xFF00 && address <= 0xFF7F:
		return m.IoPorts[address-0xFF00]
	case address >= 0xFF80 && address <= 0xFFFE:
		return m.HRAM[address-0xFF80]
	case address == 0xFFFF:
		return m.InterruptEnable
	default:
		fmt.Println("Error ocurred trying to read memory address %v", address)
	}
	return 0
}

func (m *mmu) WriteByte(address uint16, value uint8) {
	switch {
	case address >= 0x0000 && address <= 0x7FFF:
		m.ROM[address] = value
	case address >= 0x8000 && address <= 0x9FFF:
		/*
			if !m.CanAccessVRAM() {
				fmt.Println("Accessing VRAM at an illegal time...")
			}
		*/
		m.VRAM[address-0x8000] = value
	case address >= 0xA000 && address <= 0xBFFF:
		m.SwitchableRAM[address-0xA000] = value
	case address >= 0xC000 && address <= 0xDFFF:
		m.InternalRAM[address-0xC000] = value
		// Echo RAM contains the same values as internal RAM
		m.EchoRAM[address-0xC000] = value
	case address >= 0xE000 && address <= 0xFDFF:
		m.EchoRAM[address-0xE000] = value
	case address >= 0xFE00 && address <= 0xFE9F:
		/*
			if !m.CanAccessOAM() {
				fmt.Println("Accessing OAM at an illegal time...")
				return
			}
		*/
		m.OAM[address-0xFE00] = value
	case address >= 0xFF00 && address <= 0xFF7F:
		if address == 0xFF04 { // Writing to the divider register always resets it to 0
			value = 0
		} else if address == 0xFF46 {
			m.startDMA(value)
		}
		m.IoPorts[address-0xFF00] = value
	case address >= 0xFF80 && address <= 0xFFFE:
		m.HRAM[address-0xFF80] = value
	case address == 0xFFFF:
		m.InterruptEnable = value
	default:
		fmt.Println("Error ocurred trying to read memory address %v", address)
	}
}

func (m *mmu) LCDStatusMode() uint8 {
	return m.ReadAt(0xFF41) & 0x03
}

func (m *mmu) SetLCDStatusMode(mode uint8) {
	value := m.ReadAt(0xFF41)&0xFC + mode
	m.WriteByte(0xFF41, value)
}

func (m *mmu) CanAccessOAM() bool {
	mode := m.LCDStatusMode()
	return (mode == 0 || mode == 1)
}

func (m *mmu) CanAccessVRAM() bool {
	mode := m.LCDStatusMode()
	return mode != 3
}

func (m *mmu) LCDEnabled() bool {
	return GetBit(m.ReadAt(0xFF40), 7) == 1
}

func (m *mmu) WindowTileMap() uint16 {
	if GetBit(m.ReadAt(0xFF40), 6) == 1 {
		return uint16(0x9C00)
	} else {
		return uint16(0x9800)
	}
}

func (m *mmu) WindowDisplayEnabled() bool {
	return GetBit(m.ReadAt(0xFF40), 5) == 1
}

func (m *mmu) BGAndWindowAddressMode() AddressMode {
	if GetBit(m.ReadAt(0xFF40), 4) == 1 {
		return ADDRESS_MODE_8000
	} else {
		return ADDRESS_MODE_8800
	}
}

func (m *mmu) BGTileMap() uint16 {
	if GetBit(m.ReadAt(0xFF40), 3) == 1 {
		return uint16(0x9C00)
	} else {
		return uint16(0x9800)
	}
}

func (m *mmu) SpriteSize() int {
	return GetBit(m.ReadAt(0xFF40), 2)
}

func (m *mmu) SpritesEnabled() bool {
	return GetBit(m.ReadAt(0xFF40), 1) == 1
}

func (m *mmu) BGDisplayPriority() bool {
	return GetBit(m.ReadAt(0xFF40), 0) == 1
}

func (m *mmu) Tick() {
}

func (m *mmu) bgShadeForColor0() RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF47), 1)
	lowBit := GetBit(m.ReadAt(0xFF47), 0)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) bgShadeForColor1() RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF47), 3)
	lowBit := GetBit(m.ReadAt(0xFF47), 2)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) bgShadeForColor2() RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF47), 5)
	lowBit := GetBit(m.ReadAt(0xFF47), 4)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) bgShadeForColor3() RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF47), 7)
	lowBit := GetBit(m.ReadAt(0xFF47), 6)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) ConvertNumToPixel(i int) RGBPixel {
	switch i {
	case 0:
		return m.bgShadeForColor0()
	case 1:
		return m.bgShadeForColor1()
	case 2:
		return m.bgShadeForColor2()
	default:
		return m.bgShadeForColor3()
	}
}

func (m *mmu) DisableInterrupts() {
	m.WriteByte(0xFFFF, 0)
}

func (m *mmu) HasPendingInterrupt() bool {
	return m.ReadAt(0xFFFF)&m.ReadAt(0xFF0F) != 0
}

func (m *mmu) GetNextPendingInterrupt() uint16 {
	return m.interruptMapping[GetHighestBit(m.ReadAt(0xFFFF)&m.ReadAt(0xFF0F))]
}

func (m *mmu) ClearHighestInterrupt() {
	interruptFlags := m.ReadAt(0xFF0F)
	bit := GetHighestBit(m.ReadAt(0xFFFF) & interruptFlags)
	temp := (1 << uint(bit))
	m.WriteByte(0xFF0F, uint8(temp^31))
}

func (m *mmu) FireInterrupt(interrupt Interrupt) {
	val := m.ReadAt(0xFF0F) | (1 << uint(interrupt))
	m.WriteByte(0xFF0F, val)
}

func (m *mmu) startDMA(value uint8) {
	addr := uint16(value) << 8
	for i := 0; i < 0xA0; i++ {
		m.WriteByte(0xFE00+uint16(i), m.ReadAt(addr+uint16(i)))
	}
}
