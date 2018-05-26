package main

type EightBitRegister byte
type SixteenBitRegister uint16

type CPU struct {

	PC   SixteenBitRegister // Should be initialized to 0x100 to start execution
	SP	 SixteenBitRegister // Should be initialized to 0xFFFE on startup (grows downward in RAM)

	RegA EightBitRegister
	RegB EightBitRegister
	RegC EightBitRegister
	RegD EightBitRegister
	RegE EightBitRegister
	RegF EightBitRegister
	RegH EightBitRegister
	RegL EightBitRegister
}

func (c *CPU) reset() {
	c.PC = 0x100
	c.SP = 0xFFFE
}
