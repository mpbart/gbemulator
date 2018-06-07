package main

type Parameters []byte

type Instruction interface {
	Execute(Parameters)
	GetNumParameterBytes() int
}

type loadRegisterWithOffsetInstruction struct {
	cycles     int
	paramBytes int
	dest       Register
	memAddr    uint16
	offsetReg  Register
	regs       Registers
}

type loadImmediateInstruction struct {
	cycles     int
	paramBytes int
	dest       Register
	regs       Registers
}

type loadTwoByteImmediateInstruction struct {
	cycles     int
	paramBytes int
	dest1      Register
	dest2      Register
	regs       Registers
}

type loadRegisterInstruction struct {
	cycles     int
	paramBytes int
	dest       Register
	source     Register
	regs       Registers
}

type loadMemoryWithRegisterInstruction struct {
	cycles     int
	paramBytes int
	memAddr    uint16
	offset     Register
	source     Register
}

type noopInstruction struct {
	cycles     int
	paramBytes int
}

func CreateInstructions(regs Registers, mmu MMU) map[byte]Instruction {
	// Opcodes to investigate how to handle: 0xEA, 0x3A, 0x32, 0x2A, 0x22, 0xE0, 0xF0
	// 0xF8, 0xF9, 0x08, 0xF5, 0xC5, 0xD5, 0xE5, 0xF1, 0xC1, 0xD1, 0xE1
	return map[byte]Instruction{
		0x00: &noopInstruction{1, 0},

		0x06: &loadImmediateInstruction{8, 1, b, regs},
		0x0E: &loadImmediateInstruction{8, 1, c, regs},
		0x16: &loadImmediateInstruction{8, 1, d, regs},
		0x1E: &loadImmediateInstruction{8, 1, e, regs},
		0x26: &loadImmediateInstruction{8, 1, h, regs},
		0x2E: &loadImmediateInstruction{8, 1, l, regs},

		0x78: &loadRegisterInstruction{4, 0, a, b, regs},
		0x79: &loadRegisterInstruction{4, 0, a, c, regs},
		0x7A: &loadRegisterInstruction{4, 0, a, d, regs},
		0x7B: &loadRegisterInstruction{4, 0, a, e, regs},
		0x7C: &loadRegisterInstruction{4, 0, a, h, regs},
		0x7D: &loadRegisterInstruction{4, 0, a, l, regs},
		// ??? 0x7E: &loadRegisterInstruction{8, 1, a, hl, regs},
		0x7F: &loadRegisterInstruction{4, 0, a, a, regs},
		0x40: &loadRegisterInstruction{4, 0, b, b, regs},
		0x41: &loadRegisterInstruction{4, 0, b, c, regs},
		0x42: &loadRegisterInstruction{4, 0, b, d, regs},
		0x43: &loadRegisterInstruction{4, 0, b, e, regs},
		0x44: &loadRegisterInstruction{4, 0, b, h, regs},
		0x45: &loadRegisterInstruction{4, 0, b, l, regs},
		// ??? 0x46: &loadRegisterInstruction{8, 1, b, hl, regs},
		0x47: &loadRegisterInstruction{4, 0, b, a, regs},
		0x48: &loadRegisterInstruction{4, 0, c, b, regs},
		0x49: &loadRegisterInstruction{4, 0, c, c, regs},
		0x4A: &loadRegisterInstruction{4, 0, c, d, regs},
		0x4B: &loadRegisterInstruction{4, 0, c, e, regs},
		0x4C: &loadRegisterInstruction{4, 0, c, h, regs},
		0x4D: &loadRegisterInstruction{4, 0, c, l, regs},
		// ??? 0x4E: &loadRegisterInstruction{8, 0, c, hl, regs},
		0x4F: &loadRegisterInstruction{4, 0, c, a, regs},
		0x50: &loadRegisterInstruction{4, 0, d, b, regs},
		0x51: &loadRegisterInstruction{4, 0, d, c, regs},
		0x52: &loadRegisterInstruction{4, 0, d, d, regs},
		0x53: &loadRegisterInstruction{4, 0, d, e, regs},
		0x54: &loadRegisterInstruction{4, 0, d, h, regs},
		0x55: &loadRegisterInstruction{4, 0, d, l, regs},
		// ??? 0x56: &loadRegisterInstruction{8, 0, d, hl, regs},
		0x57: &loadRegisterInstruction{4, 0, e, a, regs},
		0x58: &loadRegisterInstruction{4, 0, e, b, regs},
		0x59: &loadRegisterInstruction{4, 0, e, c, regs},
		0x5A: &loadRegisterInstruction{4, 0, e, d, regs},
		0x5B: &loadRegisterInstruction{4, 0, e, e, regs},
		0x5C: &loadRegisterInstruction{4, 0, e, h, regs},
		0x5D: &loadRegisterInstruction{4, 0, e, l, regs},
		// ??? 0x5E: &loadRegisterInstruction{8, 0, e, hl, regs},
		0x5F: &loadRegisterInstruction{4, 0, e, a, regs},
		0x60: &loadRegisterInstruction{4, 0, h, b, regs},
		0x61: &loadRegisterInstruction{4, 0, h, c, regs},
		0x62: &loadRegisterInstruction{4, 0, h, d, regs},
		0x63: &loadRegisterInstruction{4, 0, h, e, regs},
		0x64: &loadRegisterInstruction{4, 0, h, h, regs},
		0x65: &loadRegisterInstruction{4, 0, h, l, regs},
		// ??? 0x66: &loadRegisterInstruction{8, 0, h, hl, regs},
		0x67: &loadRegisterInstruction{4, 0, h, a, regs},
		0x68: &loadRegisterInstruction{4, 0, l, b, regs},
		0x69: &loadRegisterInstruction{4, 0, l, c, regs},
		0x6A: &loadRegisterInstruction{4, 0, l, d, regs},
		0x6B: &loadRegisterInstruction{4, 0, l, e, regs},
		0x6C: &loadRegisterInstruction{4, 0, l, h, regs},
		0x6D: &loadRegisterInstruction{4, 0, l, l, regs},
		// ??? 0x6E: &loadRegisterInstruction{8, 0, l, hl, regs},
		0x6F: &loadRegisterInstruction{4, 0, l, a, regs},
		// ??? 0x70: &loadRegisterInstruction{8, 0, hl, b, regs},
		// ??? 0x71: &loadRegisterInstruction{8, 0, hl, c, regs},
		// ??? 0x72: &loadRegisterInstruction{8, 0, hl, d, regs},
		// ??? 0x73: &loadRegisterInstruction{8, 0, hl, e, regs},
		// ??? 0x74: &loadRegisterInstruction{8, 0, hl, h, regs},
		// ??? 0x75: &loadRegisterInstruction{8, 0, hl, l, regs},
		// TODO: wtf is this
		// =================
		// ??? 0x36: &loadRegisterInstruction{8, 0, hl, regs, mmu},
		// ================= TODO: how do you put 16 bits into an 8 bit register?
		// ??? 0x02: &loadRegisterInstruction{8, 0, bc, a, regs},
		// ??? 0x12: &loadRegisterInstruction{8, 0, de, a, regs},
		// ??? 0x77: &loadRegisterInstruction{8, 0, hl, a, regs},
		// ??? 0x0A: &loadRegisterInstruction{8, 0, a, bc, regs},
		// ??? 0x1A: &loadRegisterInstruction{8, 0, a, de, regs},

		0xF2: &loadRegisterWithOffsetInstruction{8, 0, a, 0xFF00, c, regs},
		0xE2: &loadMemoryWithRegisterInstruction{8, 0, 0xFF00, c, a},

		0x01: &loadTwoByteImmediateInstruction{12, 2, b, c, regs},
		0x11: &loadTwoByteImmediateInstruction{12, 2, d, e, regs},
		0x21: &loadTwoByteImmediateInstruction{12, 2, h, l, regs},
		// TODO: This is really hacky, fix it
		// 0x31: &loadTwoByteImmediateInstruction{12, 2, sp, sp, regs},
	}
}

// Assuming for now that loading of 2 byte immediate values is handled in different type
func (i *loadImmediateInstruction) Execute(params Parameters) {
	i.regs.WriteRegister(i.dest, params[0])
}

func (n *noopInstruction) Execute(_ Parameters) {
}

func (n *noopInstruction) GetNumParameterBytes() int {
	return n.paramBytes
}

func (i *loadImmediateInstruction) GetNumParameterBytes() int {
	return i.paramBytes
}

func (i *loadRegisterInstruction) GetNumParameterBytes() int {
	return i.paramBytes
}

func (i *loadRegisterWithOffsetInstruction) GetNumParameterBytes() int {
	return i.paramBytes
}

func (i *loadRegisterInstruction) Execute(params Parameters) {
}

func (i *loadMemoryWithRegisterInstruction) GetNumParameterBytes() int {
	return i.paramBytes
}

func (i *loadRegisterWithOffsetInstruction) Execute(params Parameters) {
}

func (i *loadMemoryWithRegisterInstruction) Execute(params Parameters) {
}

func (i *loadTwoByteImmediateInstruction) Execute(params Parameters) {
}

func (i *loadTwoByteImmediateInstruction) GetNumParameterBytes() int {
	return i.paramBytes
}
