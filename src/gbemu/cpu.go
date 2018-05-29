package main

type EightBitRegister byte
type SixteenBitRegister uint16

type CPU interface {
	Reset()
}

type cpu struct {
	mmu MMU

	PC SixteenBitRegister // Should be initialized to 0x100 to start execution
	SP SixteenBitRegister // Should be initialized to 0xFFFE on startup (grows downward in RAM)

	RegA EightBitRegister // Accumulator
	RegB EightBitRegister
	RegC EightBitRegister
	RegD EightBitRegister
	RegE EightBitRegister
	RegF EightBitRegister // Flags
	RegH EightBitRegister
	RegL EightBitRegister
}

func CreateCPU(mmu MMU) CPU {
	return &cpu{
		mmu: mmu,
	}
}

func (c *cpu) Reset() {
	c.PC = 0x100
	c.SP = 0xFFFE

	//TODO: Set register values
	c.mmu.WriteByte(0xFF05, 0x00)
	c.mmu.WriteByte(0xFF06, 0x00)
	c.mmu.WriteByte(0xFF07, 0x00)
	c.mmu.WriteByte(0xFF10, 0x80)
	c.mmu.WriteByte(0xFF11, 0xBF)
	c.mmu.WriteByte(0xFF12, 0xF3)
	c.mmu.WriteByte(0xFF14, 0xBF)
	c.mmu.WriteByte(0xFF16, 0x3F)
	c.mmu.WriteByte(0xFF17, 0x00)
	c.mmu.WriteByte(0xFF19, 0xBF)
	c.mmu.WriteByte(0xFF1A, 0x7F)
	c.mmu.WriteByte(0xFF1B, 0xFF)
	c.mmu.WriteByte(0xFF1C, 0x9F)
	c.mmu.WriteByte(0xFF1E, 0xBF)
	c.mmu.WriteByte(0xFF20, 0xFF)
	c.mmu.WriteByte(0xFF21, 0x00)
	c.mmu.WriteByte(0xFF22, 0x00)
	c.mmu.WriteByte(0xFF23, 0xBF)
	c.mmu.WriteByte(0xFF24, 0x77)
	c.mmu.WriteByte(0xFF25, 0xF3)
	c.mmu.WriteByte(0xFF26, 0xF1)
	c.mmu.WriteByte(0xFF40, 0x91)
	c.mmu.WriteByte(0xFF42, 0x00)
	c.mmu.WriteByte(0xFF43, 0x00)
	c.mmu.WriteByte(0xFF45, 0x00)
	c.mmu.WriteByte(0xFF47, 0xFC)
	c.mmu.WriteByte(0xFF48, 0xFF)
	c.mmu.WriteByte(0xFF49, 0xFF)
	c.mmu.WriteByte(0xFF4A, 0x00)
	c.mmu.WriteByte(0xFF4B, 0x00)
	c.mmu.WriteByte(0xFF50, 0x00)
	c.mmu.WriteByte(0xFFFF, 0x00)
}
