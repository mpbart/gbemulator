package main

type Loader interface {
	GetValue() byte
	GetTwoByteValue() uint16
}

type Instruction interface {
	Execute()
}

type loadImmediateInstruction struct {
	cycles int
	paramBytes int
	dest   Register
	source byte
}

type noopInstruction struct {
	cycles int
	paramBytes int
}

// TODO: Design interface so that different types of commands can be returned as Instruction type
func CreateInstructions() map[byte]Instruction {
	return map[byte]Instruction{
	0x00: &noopInstruction{1, 0},
	//0x06: &loadImmediateInstruction{8, 1, b, 0},
	//&loadInstruction{0x0E, 8, 0},
	//&loadInstruction{0x16, 8, 0},
	//&loadInstruction{0x1E, 8, 0},
	//&loadInstruction{0x26, 8, 0},
	//&loadInstruction{0x2E, 8, 0},
	}
}

func (i *loadImmediateInstruction) Execute() {
	//registers.WriteRegister(i.dest, i.source)
}

func (n *noopInstruction) Execute() {
}
