package main

type Command int

type Instructions struct {
	opcode byte
	cycles int
	command Command
}

// TODO: Design interface so that different types of commands can be returned as Instruction type
func CreateInstructions() []Instructions {
	return []Instructions{
		Instruction{0x06, 8, 0},
		Instruction{0x0E, 8, 0},
		Instruction{0x16, 8, 0},
		Instruction{0x1E, 8, 0},
		Instruction{0x26, 8, 0},
		Instruction{0x2E, 8, 0},
	}
}
