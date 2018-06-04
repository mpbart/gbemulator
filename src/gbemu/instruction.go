package main

type Loader interface {
	GetValue() byte
	GetTwoByteValue() uint16
}

type Instruction interface {
	Execute()
}

type loadInstruction struct {
	opcode byte
	cycles int
	source Loader
	dest   Loader
}

// TODO: Design interface so that different types of commands can be returned as Instruction type
func CreateInstructions() []Instruction {
	return []Instruction{
	//&loadInstruction{0x06, 8, 0, 0},
	//&loadInstruction{0x0E, 8, 0},
	//&loadInstruction{0x16, 8, 0},
	//&loadInstruction{0x1E, 8, 0},
	//&loadInstruction{0x26, 8, 0},
	//&loadInstruction{0x2E, 8, 0},
	}
}

func (i *loadInstruction) Execute() {
}
