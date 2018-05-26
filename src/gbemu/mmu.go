package main

import (
	"fmt"
)

type MMU struct {
	ROM				[32768]uint8// 0x0000 - 0x7FFF
	VRAM			[8192]uint8 // 0x8000 - 0x9FFF
	SwitchableRAM	[8192]uint8 // 0xA000 - 0xBFFF
	InternalRAM		[8192]uint8 // 0xC000 - 0xDFFF
	EchoRAM			[8192]uint8 // 0xE000 - 0xFDFF	
	OAM				[160]uint8	// 0xFE00 - 0xFE9F
	OptionalIo		[95]uint8   // 0xFEA0 - 0xFEFF
	IoPorts			[76]uint8   // 0xFF00 - 0xFF4B
	OptionalIo2		[52]uint8   // 0xFF4C - 0xFF7F
	ExtraRAM		[127]uint8  // 0xFF80 - 0xFFFE
	InterruptEnable uint8       // 0xFFFF
}

func (m *MMU) reset() {
	fmt.Println(m.ReadAt(0xFFFF))
}

func (m *MMU) ReadAt(address uint16) uint8 {
	switch {
	case address >= 0x0000 && address <= 0x7FFF:
		return m.ROM[address]
	case address >= 0x8000 && address <= 0x9FFF:
		return m.VRAM[address - 0x8000]
	case address >= 0xA000 && address <= 0xBFFF:
		return m.SwitchableRAM[address - 0xA000]
	case address >= 0xC000 && address <= 0xDFFF:
		return m.InternalRAM[address - 0xC000]
	case address >= 0xE000 && address <= 0xFDFF:
		return m.EchoRAM[address - 0xE000]
	case address >= 0xFE00 && address <= 0xFE9F:
		return m.OAM[address - 0xFE00]
	case address >= 0xFEA0 && address <= 0xFE9F:
		return m.OptionalIo[address - 0xFEA0]
	case address >= 0xFF00 && address <= 0xFF4B:
		return m.IoPorts[address - 0xFF00]
	case address >= 0xFF4C && address <= 0xFF7F:
		return m.OptionalIo2[address - 0xFF4C]
	case address >= 0xFF80 && address <= 0xFFFE:
		return m.ExtraRAM[address - 0xFF80]
	case address == 0xFFFF:
		return m.InterruptEnable
	default:
		fmt.Println("Error ocurred trying to read memory address %v", address)
	}
	return 0
}
