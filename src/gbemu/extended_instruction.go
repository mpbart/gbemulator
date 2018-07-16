package main

type ExtendedInstruction interface {
	Execute(Parameters) Addresser
	GetNumParameterBytes() int
}

type swapInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type swapFromMemoryInstruction struct {
	basicInstruction
	regs Registers
	mmu  MMU
}

func CreateExtendedInstructions(regs Registers, mmu MMU) map[byte]ExtendedInstruction {
	return map[byte]ExtendedInstruction{
		0x30: &swapInstruction{basicInstruction{8, 0}, b, regs},
		0x31: &swapInstruction{basicInstruction{8, 0}, c, regs},
		0x32: &swapInstruction{basicInstruction{8, 0}, d, regs},
		0x33: &swapInstruction{basicInstruction{8, 0}, e, regs},
		0x34: &swapInstruction{basicInstruction{8, 0}, h, regs},
		0x35: &swapInstruction{basicInstruction{8, 0}, l, regs},
		0x36: &swapFromMemoryInstruction{basicInstruction{8, 0}, regs, mmu},
		0x37: &swapInstruction{basicInstruction{8, 0}, a, regs},
	}
}

func (i *swapInstruction) Execute(params Parameters) Addresser {
	return &address{}
}

func (i *swapFromMemoryInstruction) Execute(params Parameters) Addresser {
	return &address{}
}
