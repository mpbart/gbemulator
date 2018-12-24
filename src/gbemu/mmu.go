package main

import (
	"fmt"
	"runtime/debug"
)

// I/O Registers:
// 0xFF00 - P1   - Joypad input
// 0xFF01 - SB   - Serial transfer data
// 0xFF02 - SC   - Serial I/O control
// 0xFF04 - DIV  - Divider register
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
const (
	JOYPAD_INPUT                    uint16 = 0xFF00
	SERIAL_TRANSFER_DATA            uint16 = 0xFF01
	SERIAL_IO_CONTROL               uint16 = 0xFF02
	DIVIDER_REGISTER                uint16 = 0xFF04
	TIMER_REGISTER                  uint16 = 0xFF05
	TIMER_MODULO                    uint16 = 0xFF06
	TIMER_CONTROL                   uint16 = 0xFF07
	INTERRUPT_FLAGS                 uint16 = 0xFF0F
	SWEEP_REGISTER                  uint16 = 0xFF10
	SOUND_MODE_1_WAVE_PATTERN       uint16 = 0xFF11
	SOUND_MODE_1_ENVELOPE           uint16 = 0xFF12
	SOUND_MODE_1_FREQUENCY_LOW      uint16 = 0xFF13
	SOUND_MODE_1_FREQUENCY_HIGH     uint16 = 0xFF14
	SOUND_MODE_2_WAVE_PATTERN       uint16 = 0xFF16
	SOUND_MODE_2_ENVELOPE           uint16 = 0xFF17
	SOUND_MODE_2_FREQUENCY_LOW      uint16 = 0xFF18
	SOUND_MODE_2_FREQUENCY_HIGH     uint16 = 0xFF19
	SOUND_MODE_3_ON_OFF             uint16 = 0xFF1A
	SOUND_MODE_3_WAVE_PATTERN       uint16 = 0xFF1B
	SOUND_MODE_3_OUTPUT_SELECT      uint16 = 0xFF1C
	SOUND_MODE_3_FREQUENCY_LOW      uint16 = 0xFF1D
	SOUND_MODE_3_FREQUENCY_HIGH     uint16 = 0xFF1E
	SOUND_MODE_4_WAVE_PATTERN       uint16 = 0xFF20
	SOUND_MODE_4_ENVELOPE           uint16 = 0xFF21
	SOUND_MODE_4_POLYNOMIAL_COUNTER uint16 = 0xFF22
	SOUND_MODE_4_COUNTER            uint16 = 0xFF23
	CHANNEL_CONTROL                 uint16 = 0xFF24
	SOUND_OUTPUT_SELECTION          uint16 = 0xFF25
	SOUND_ON_OFF                    uint16 = 0xFF26
	// WAVE_PATTERN_RAM                uint16 = 0xFF30 - 0xFF3F
	LCD_CONTROL             uint16 = 0xFF40
	LCDC_STATUS             uint16 = 0xFF41
	SCROLL_Y                uint16 = 0xFF42
	SCROLL_X                uint16 = 0xFF43
	LCDC_Y_COORDINATE       uint16 = 0xFF44
	LY_COMPARE              uint16 = 0xFF45
	DMA_TRANSFER_ADDRESS    uint16 = 0xFF46
	BG_WINDOW_PALLETTE_DATA uint16 = 0xFF47
	OBJECT_PALLETTE_0       uint16 = 0xFF48
	OBJECT_PALLETTE_1       uint16 = 0xFF49
	WINDOW_Y_POSITION       uint16 = 0xFF4A
	WINDOW_X_POSITION       uint16 = 0xFF4B
	INTERRUPT_ENABLE        uint16 = 0xFFFF
)

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
	BGDisplayEnabled() bool
	ConvertNumToBgPixel(int) RGBPixel
	ConvertNumToSpritePixel(int, int) RGBPixel
	HasPendingInterrupt() bool
	GetNextPendingInterrupt() uint16
	ClearHighestInterrupt()
	FireInterrupt(Interrupt)
	ReadJoypadInput() uint8
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
	// Initialize memory from values specified in manual
	m.WriteByte(TIMER_REGISTER, 0x00)
	m.WriteByte(TIMER_MODULO, 0x00)
	m.WriteByte(TIMER_CONTROL, 0x00)
	m.WriteByte(LCD_CONTROL, 0x91)
	m.WriteByte(SCROLL_Y, 0x00)
	m.WriteByte(SCROLL_X, 0x00)
	m.WriteByte(LY_COMPARE, 0x00)
	m.WriteByte(BG_WINDOW_PALLETTE_DATA, 0xFC)
	m.WriteByte(OBJECT_PALLETTE_0, 0xFF)
	m.WriteByte(OBJECT_PALLETTE_1, 0xFF)
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
		switch address {
		case JOYPAD_INPUT:
			return m.ReadJoypadInput()
		}
		return m.IoPorts[address-0xFF00]
	case address >= 0xFF80 && address <= 0xFFFE:
		return m.HRAM[address-0xFF80]
	case address == INTERRUPT_ENABLE:
		return m.InterruptEnable
	default:
		fmt.Printf("Error ocurred trying to read memory address %x\n", address)
		debug.PrintStack()
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
	case address >= 0xFEA0 && address <= 0xFEFF:
		return
	case address >= 0xFF00 && address <= 0xFF7F:
		switch address {
		case DIVIDER_REGISTER: // Writing to the divider register always resets it to 0
			value = 0
		case DMA_TRANSFER_ADDRESS:
			m.startDMA(value)
		case LCDC_STATUS:
			value = (value & 0x78) | (m.IoPorts[0x41] & 0x87) // Bits 8, 2, 1, 0 are read-only
		case JOYPAD_INPUT:
			value &= 0x30 // Only bits 4 and 5 can be set
		default:
		}
		m.IoPorts[address-0xFF00] = value
	case address >= 0xFF80 && address <= 0xFFFE:
		fmt.Printf("Writing %x to %x\n", value, address)
		m.HRAM[address-0xFF80] = value
	case address == INTERRUPT_ENABLE:
		m.InterruptEnable = value
	default:
		fmt.Printf("Error ocurred trying to write memory address %x\n", address)
		debug.PrintStack()
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

func (m *mmu) BGDisplayEnabled() bool {
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

func (m *mmu) spriteShadeForColor0(paletteNum int) RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 1)
	lowBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 0)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) spriteShadeForColor1(paletteNum int) RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 3)
	lowBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 2)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) spriteShadeForColor2(paletteNum int) RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 5)
	lowBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 4)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) spriteShadeForColor3(paletteNum int) RGBPixel {
	highBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 7)
	lowBit := GetBit(m.ReadAt(0xFF48+uint16(paletteNum)), 6)
	return m.colorMapping[BitsToNum(highBit, lowBit)]
}

func (m *mmu) ConvertNumToBgPixel(i int) RGBPixel {
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

func (m *mmu) ConvertNumToSpritePixel(i, paletteNum int) RGBPixel {
	switch i {
	case 0:
		return m.spriteShadeForColor0(paletteNum)
	case 1:
		return m.spriteShadeForColor1(paletteNum)
	case 2:
		return m.spriteShadeForColor2(paletteNum)
	default:
		return m.spriteShadeForColor3(paletteNum)
	}
}

func (m *mmu) HasPendingInterrupt() bool {
	return m.ReadAt(INTERRUPT_ENABLE)&m.ReadAt(INTERRUPT_FLAGS) != 0
}

func (m *mmu) GetNextPendingInterrupt() uint16 {
	return m.interruptMapping[GetHighestInterruptBit(m.ReadAt(INTERRUPT_ENABLE)&m.ReadAt(INTERRUPT_FLAGS))]
}

func (m *mmu) ClearHighestInterrupt() {
	interruptFlags := m.ReadAt(INTERRUPT_FLAGS)
	bit := GetHighestInterruptBit(m.ReadAt(INTERRUPT_ENABLE) & interruptFlags)
	temp := (1 << uint(bit))
	m.WriteByte(INTERRUPT_FLAGS, interruptFlags&uint8(temp^31))
}

func (m *mmu) FireInterrupt(interrupt Interrupt) {
	val := m.ReadAt(INTERRUPT_FLAGS) | (1 << uint(interrupt))
	m.WriteByte(INTERRUPT_FLAGS, val)
}

func (m *mmu) startDMA(value uint8) {
	addr := uint16(value) << 8
	for i := 0; i < 0xA0; i++ {
		m.WriteByte(0xFE00+uint16(i), m.ReadAt(addr+uint16(i)))
	}
}

func (m *mmu) ReadJoypadInput() uint8 {
	if GetBit(m.ReadAt(JOYPAD_INPUT), 4) == 0 { // Get direction key inputs
		// bit 3 - down
		// bit 2 - up
		// bit 1 - left
		// bit 0 - right
		return 0xEF
	} else { // Get button keys input
		// bit 3 - start
		// bit 2 - select
		// bit 1 - B
		// bit 0 - A
		return 0xDF
	}
}
