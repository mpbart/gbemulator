package main

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
}
