package main

import (
	"fmt"
)

type MMU interface {
	Reset()
	ReadAt(uint16) uint8
	WriteByte(uint16, uint8)
}

type mmu struct {
	ROM             [32768]uint8 // 0x0000 - 0x7FFF
	VRAM            [8192]uint8  // 0x8000 - 0x9FFF
	SwitchableRAM   [8192]uint8  // 0xA000 - 0xBFFF
	InternalRAM     [8192]uint8  // 0xC000 - 0xDFFF
	EchoRAM         [8192]uint8  // 0xE000 - 0xFDFF
	OAM             [160]uint8   // 0xFE00 - 0xFE9F
	Unused          [95]uint8    // 0xFEA0 - 0xFEFF
	IoPorts         [128]uint8   // 0xFF00 - 0xFF7F
	HRAM            [127]uint8   // 0xFF80 - 0xFFFE
	InterruptEnable uint8        // 0xFFFF
}

func CreateMMU() MMU {
	return &mmu{}
}

func (m *mmu) Reset() {
	fmt.Println(m.ReadAt(0xFFFF))
}

func (m *mmu) ReadAt(address uint16) uint8 {
	switch {
	case address >= 0x0000 && address <= 0x7FFF:
		return m.ROM[address]
	case address >= 0x8000 && address <= 0x9FFF:
		if !m.CanAccessVRAM() {
			fmt.Println("Accessing VRAM at an illegal time...")
			return 0
		}
		return m.VRAM[address-0x8000]
	case address >= 0xA000 && address <= 0xBFFF:
		return m.SwitchableRAM[address-0xA000]
	case address >= 0xC000 && address <= 0xDFFF:
		return m.InternalRAM[address-0xC000]
	case address >= 0xE000 && address <= 0xFDFF:
		return m.EchoRAM[address-0xE000]
	case address >= 0xFE00 && address <= 0xFE9F:
		if !m.CanAccessOAM() {
			fmt.Println("Accessing OAM at an illegal time...")
			return 0
		}
		return m.OAM[address-0xFE00]
	// TODO: Check for accessing unused memory locations and panic?
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
		if !m.CanAccessVRAM() {
			fmt.Println("Accessing VRAM at an illegal time...")
			return
		}
		m.VRAM[address-0x8000] = value
	case address >= 0xA000 && address <= 0xBFFF:
		m.SwitchableRAM[address-0xA000] = value
	case address >= 0xC000 && address <= 0xDFFF:
		m.InternalRAM[address-0xC000] = value
	case address >= 0xE000 && address <= 0xFDFF:
		m.EchoRAM[address-0xE000] = value
	case address >= 0xFE00 && address <= 0xFE9F:
		if !m.CanAccessOAM() {
			fmt.Println("Accessing OAM at an illegal time...")
			return
		}
		m.OAM[address-0xFE00] = value
	// TODO: Check for accessing unused memory locations and panic?
	case address >= 0xFF00 && address <= 0xFF7F:
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

func (m *mmu) CanAccessOAM() bool {
	mode := m.LCDStatusMode()
	return (mode == 0 || mode == 1)
}

func (m *mmu) CanAccessVRAM() bool {
	mode := m.LCDStatusMode()
	return mode != 3
}
