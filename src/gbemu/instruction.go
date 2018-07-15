package main

import "fmt"

type Parameters []byte

type Addresser interface {
	ShouldJump() bool
	NewAddress() uint16
}

type address struct {
	pcShouldJump bool
	newAddress   uint16
}

type Instruction interface {
	Execute(Parameters) Addresser
	GetNumParameterBytes() int
}

type basicInstruction struct {
	cycles     int
	paramBytes int
}

type loadRegisterWithOffsetInstruction struct {
	basicInstruction
	dest      Register
	memAddr   uint16
	offsetReg Register
	regs      Registers
	mmu       MMU
}

type loadImmediateInstruction struct {
	basicInstruction
	dest Register
	regs Registers
}

type loadTwoByteImmediateInstruction struct {
	basicInstruction
	dest1 Register
	dest2 Register
	regs  Registers
}

type loadRegisterInstruction struct {
	basicInstruction
	dest   Register
	source Register
	regs   Registers
}

type loadMemoryWithRegisterInstruction struct {
	basicInstruction
	memAddr uint16
	offset  Register
	source  Register
	regs    Registers
	mmu     MMU
}

type noopInstruction struct {
	basicInstruction
}

type pushInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
}

type popInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
}

type addInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type addCarryInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type addCarryFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	dest    Register
	regs    Registers
	mmu     MMU
}

type addImmediateInstruction struct {
	basicInstruction
	regs Registers
}

type subInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type subCarryInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type andInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type orInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type xorInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type cpInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type incInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type decInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type add16BitInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
}

type add16BitFromSPInstruction struct {
	basicInstruction
	regs Registers
}

type addSP16BitInstruction struct {
	basicInstruction
	regs Registers
}

type inc16BitInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
}

type incSP16BitInstruction struct {
	basicInstruction
	regs Registers
}

type dec16BitInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
}

type decSP16BitInstruction struct {
	basicInstruction
	regs Registers
}

type swapInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type daaInstruction struct {
	basicInstruction
	regs Registers
}

type cplInstruction struct {
	basicInstruction
	regs Registers
}

type ccfInstruction struct {
	basicInstruction
	regs Registers
}

type scfInstruction struct {
	basicInstruction
	regs Registers
}

type haltInstruction struct {
	basicInstruction
}

type stopInstruction struct {
	basicInstruction
}

type diInstruction struct {
	basicInstruction
}

type eiInstruction struct {
	basicInstruction
}

type jumpInstruction struct {
	basicInstruction
}

type conditionalJumpInstruction struct {
	basicInstruction
	regs        Registers
	conditional func() bool
}

type jumpHlInstruction struct {
	basicInstruction
	regs Registers
}

type jumpImmediateInstruction struct {
	basicInstruction
	regs Registers
}

type conditionalJumpImmediateInstruction struct {
	basicInstruction
	regs        Registers
	conditional func() bool
}

type callInstruction struct {
	basicInstruction
	regs Registers
}

type restartInstruction struct {
	basicInstruction
	regs   Registers
	offset uint16
}

type returnInstruction struct {
	basicInstruction
	regs Registers
}

type callConditionalInstruction struct {
	basicInstruction
	regs        Registers
	conditional func() bool
}

type returnConditionalInstruction struct {
	basicInstruction
	regs        Registers
	conditional func() bool
}

type retiInstruction struct {
	basicInstruction
	regs Registers
}

type writeMemoryInstruction struct {
	basicInstruction
	dest1  Register
	dest2  Register
	source Register
	regs   Registers
	mmu    MMU
}

type loadRegisterFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	dest    Register
	regs    Registers
	mmu     MMU
}

type addFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type subFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type subCarryFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type andFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type addCarryImmediateInstruction struct {
	basicInstruction
	regs Registers
}

type orFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type lddInstruction struct {
	basicInstruction
	dest    Register
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type lddhlInstruction struct {
	basicInstruction
	dest1  Register
	dest2  Register
	source Register
	regs   Registers
	mmu    MMU
}

type ldiInstruction struct {
	basicInstruction
	dest    Register
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type ldihlInstruction struct {
	basicInstruction
	dest1  Register
	dest2  Register
	source Register
	regs   Registers
	mmu    MMU
}

type ldhImmediateInstruction struct {
	basicInstruction
	source Register
	regs   Registers
	mmu    MMU
}

type ldhaInstruction struct {
	basicInstruction
	dest Register
	regs Registers
	mmu  MMU
}

type ldsphlInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
}

type ldhlspInstruction struct {
	basicInstruction
	dest1 Register
	dest2 Register
	regs  Registers
}

func CreateInstructions(regs Registers, mmu MMU) map[byte]Instruction {
	// Opcodes to investigate how to handle: 0xEA,
	// 0xF8, 0x08, 0xF5, 0xC5, 0xD5, 0xE5, 0xF1, 0xC1, 0xD1, 0xE1, 0xE8 0xCBXX ?? WTF?!?
	// 0x07, 0x17, 0x0F, 0x1F
	return map[byte]Instruction{
		0x00: &noopInstruction{basicInstruction{1, 0}},

		0x06: &loadImmediateInstruction{basicInstruction{8, 1}, b, regs},
		0x0E: &loadImmediateInstruction{basicInstruction{8, 1}, c, regs},
		0x16: &loadImmediateInstruction{basicInstruction{8, 1}, d, regs},
		0x1E: &loadImmediateInstruction{basicInstruction{8, 1}, e, regs},
		0x26: &loadImmediateInstruction{basicInstruction{8, 1}, h, regs},
		0x2E: &loadImmediateInstruction{basicInstruction{8, 1}, l, regs},

		0x78: &loadRegisterInstruction{basicInstruction{4, 0}, a, b, regs},
		0x79: &loadRegisterInstruction{basicInstruction{4, 0}, a, c, regs},
		0x7A: &loadRegisterInstruction{basicInstruction{4, 0}, a, d, regs},
		0x7B: &loadRegisterInstruction{basicInstruction{4, 0}, a, e, regs},
		0x7C: &loadRegisterInstruction{basicInstruction{4, 0}, a, h, regs},
		0x7D: &loadRegisterInstruction{basicInstruction{4, 0}, a, l, regs},
		0x7F: &loadRegisterInstruction{basicInstruction{4, 0}, a, a, regs},
		0x22: &ldihlInstruction{basicInstruction{8, 0}, h, l, a, regs, mmu},
		0x2A: &ldiInstruction{basicInstruction{8, 0}, a, h, l, regs, mmu},
		0x32: &lddhlInstruction{basicInstruction{8, 0}, h, l, a, regs, mmu},
		0x3A: &lddInstruction{basicInstruction{8, 0}, a, h, l, regs, mmu},
		0x40: &loadRegisterInstruction{basicInstruction{4, 0}, b, b, regs},
		0x41: &loadRegisterInstruction{basicInstruction{4, 0}, b, c, regs},
		0x42: &loadRegisterInstruction{basicInstruction{4, 0}, b, d, regs},
		0x43: &loadRegisterInstruction{basicInstruction{4, 0}, b, e, regs},
		0x44: &loadRegisterInstruction{basicInstruction{4, 0}, b, h, regs},
		0x45: &loadRegisterInstruction{basicInstruction{4, 0}, b, l, regs},
		0x46: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, h, l, b, regs, mmu},
		0x47: &loadRegisterInstruction{basicInstruction{4, 0}, b, a, regs},
		0x48: &loadRegisterInstruction{basicInstruction{4, 0}, c, b, regs},
		0x49: &loadRegisterInstruction{basicInstruction{4, 0}, c, c, regs},
		0x4A: &loadRegisterInstruction{basicInstruction{4, 0}, c, d, regs},
		0x4B: &loadRegisterInstruction{basicInstruction{4, 0}, c, e, regs},
		0x4C: &loadRegisterInstruction{basicInstruction{4, 0}, c, h, regs},
		0x4D: &loadRegisterInstruction{basicInstruction{4, 0}, c, l, regs},
		0x4E: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, h, l, c, regs, mmu},
		0x4F: &loadRegisterInstruction{basicInstruction{4, 0}, c, a, regs},
		0x50: &loadRegisterInstruction{basicInstruction{4, 0}, d, b, regs},
		0x51: &loadRegisterInstruction{basicInstruction{4, 0}, d, c, regs},
		0x52: &loadRegisterInstruction{basicInstruction{4, 0}, d, d, regs},
		0x53: &loadRegisterInstruction{basicInstruction{4, 0}, d, e, regs},
		0x54: &loadRegisterInstruction{basicInstruction{4, 0}, d, h, regs},
		0x55: &loadRegisterInstruction{basicInstruction{4, 0}, d, l, regs},
		0x56: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, h, l, d, regs, mmu},
		0x57: &loadRegisterInstruction{basicInstruction{4, 0}, e, a, regs},
		0x58: &loadRegisterInstruction{basicInstruction{4, 0}, e, b, regs},
		0x59: &loadRegisterInstruction{basicInstruction{4, 0}, e, c, regs},
		0x5A: &loadRegisterInstruction{basicInstruction{4, 0}, e, d, regs},
		0x5B: &loadRegisterInstruction{basicInstruction{4, 0}, e, e, regs},
		0x5C: &loadRegisterInstruction{basicInstruction{4, 0}, e, h, regs},
		0x5D: &loadRegisterInstruction{basicInstruction{4, 0}, e, l, regs},
		0x5E: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, h, l, e, regs, mmu},
		0x5F: &loadRegisterInstruction{basicInstruction{4, 0}, e, a, regs},
		0x60: &loadRegisterInstruction{basicInstruction{4, 0}, h, b, regs},
		0x61: &loadRegisterInstruction{basicInstruction{4, 0}, h, c, regs},
		0x62: &loadRegisterInstruction{basicInstruction{4, 0}, h, d, regs},
		0x63: &loadRegisterInstruction{basicInstruction{4, 0}, h, e, regs},
		0x64: &loadRegisterInstruction{basicInstruction{4, 0}, h, h, regs},
		0x65: &loadRegisterInstruction{basicInstruction{4, 0}, h, l, regs},
		0x66: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, h, l, h, regs, mmu},
		0x67: &loadRegisterInstruction{basicInstruction{4, 0}, h, a, regs},
		0x68: &loadRegisterInstruction{basicInstruction{4, 0}, l, b, regs},
		0x69: &loadRegisterInstruction{basicInstruction{4, 0}, l, c, regs},
		0x6A: &loadRegisterInstruction{basicInstruction{4, 0}, l, d, regs},
		0x6B: &loadRegisterInstruction{basicInstruction{4, 0}, l, e, regs},
		0x6C: &loadRegisterInstruction{basicInstruction{4, 0}, l, h, regs},
		0x6D: &loadRegisterInstruction{basicInstruction{4, 0}, l, l, regs},
		0x6E: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, h, l, l, regs, mmu},
		0x6F: &loadRegisterInstruction{basicInstruction{4, 0}, l, a, regs},
		0x70: &writeMemoryInstruction{basicInstruction{8, 0}, h, l, b, regs, mmu},
		0x71: &writeMemoryInstruction{basicInstruction{8, 0}, h, l, c, regs, mmu},
		0x72: &writeMemoryInstruction{basicInstruction{8, 0}, h, l, d, regs, mmu},
		0x73: &writeMemoryInstruction{basicInstruction{8, 0}, h, l, e, regs, mmu},
		0x74: &writeMemoryInstruction{basicInstruction{8, 0}, h, l, h, regs, mmu},
		0x75: &writeMemoryInstruction{basicInstruction{8, 0}, h, l, l, regs, mmu},
		// TODO: wtf is this
		// =================
		// ??? 0x36: &loadRegisterInstruction{8, 0, hl, regs, mmu},
		// ================= TODO: how do you put 16 bits into an 8 bit register?
		0x02: &writeMemoryInstruction{basicInstruction{8, 0}, b, c, a, regs, mmu},
		0x12: &writeMemoryInstruction{basicInstruction{8, 0}, d, e, a, regs, mmu},
		0x77: &writeMemoryInstruction{basicInstruction{8, 0}, h, l, a, regs, mmu},
		0x0A: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, b, c, a, regs, mmu},
		0x1A: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, d, e, a, regs, mmu},
		0x7E: &loadRegisterFromMemoryInstruction{basicInstruction{8, 0}, h, l, a, regs, mmu},

		0xF2: &loadRegisterWithOffsetInstruction{basicInstruction{8, 0}, a, 0xFF00, c, regs, mmu},
		0xE2: &loadMemoryWithRegisterInstruction{basicInstruction{8, 0}, 0xFF00, c, a, regs, mmu},

		0x01: &loadTwoByteImmediateInstruction{basicInstruction{12, 2}, b, c, regs},
		0x11: &loadTwoByteImmediateInstruction{basicInstruction{12, 2}, d, e, regs},
		0x21: &loadTwoByteImmediateInstruction{basicInstruction{12, 2}, h, l, regs},
		// TODO: This is really hacky, fix it
		// 0x31: &loadTwoByteImmediateInstruction{12, 2, sp, sp, regs},

		0xF5: &pushInstruction{basicInstruction{16, 0}, a, f, regs},
		0xF8: &ldhlspInstruction{basicInstruction{12, 1}, h, l, regs},
		0xF9: &ldsphlInstruction{basicInstruction{8, 0}, h, l, regs},
		0xC5: &pushInstruction{basicInstruction{16, 0}, b, c, regs},
		0xD5: &pushInstruction{basicInstruction{16, 0}, d, e, regs},
		0xE0: &ldhImmediateInstruction{basicInstruction{12, 1}, a, regs, mmu},
		0xE5: &pushInstruction{basicInstruction{16, 0}, h, l, regs},
		0xF0: &ldhaInstruction{basicInstruction{12, 1}, a, regs, mmu},
		0xF1: &popInstruction{basicInstruction{12, 0}, a, f, regs},
		0xC1: &popInstruction{basicInstruction{12, 0}, b, c, regs},
		0xD1: &popInstruction{basicInstruction{12, 0}, d, e, regs},
		0xE1: &popInstruction{basicInstruction{12, 0}, h, l, regs},
		0x87: &addInstruction{basicInstruction{4, 0}, a, regs},
		0x80: &addInstruction{basicInstruction{4, 0}, b, regs},
		0x81: &addInstruction{basicInstruction{4, 0}, c, regs},
		0x82: &addInstruction{basicInstruction{4, 0}, d, regs},
		0x83: &addInstruction{basicInstruction{4, 0}, e, regs},
		0x84: &addInstruction{basicInstruction{4, 0}, h, regs},
		0x85: &addInstruction{basicInstruction{4, 0}, l, regs},
		0x86: &addFromMemoryInstruction{basicInstruction{8, 0}, h, l, regs, mmu},
		0xC6: &addImmediateInstruction{basicInstruction{8, 1}, regs},
		0x8F: &addCarryInstruction{basicInstruction{4, 0}, a, regs},
		0x88: &addCarryInstruction{basicInstruction{4, 0}, b, regs},
		0x89: &addCarryInstruction{basicInstruction{4, 0}, c, regs},
		0x8A: &addCarryInstruction{basicInstruction{4, 0}, d, regs},
		0x8B: &addCarryInstruction{basicInstruction{4, 0}, e, regs},
		0x8C: &addCarryInstruction{basicInstruction{4, 0}, h, regs},
		0x8D: &addCarryInstruction{basicInstruction{4, 0}, l, regs},
		0x8E: &addCarryFromMemoryInstruction{basicInstruction{8, 0}, h, l, a, regs, mmu},
		0xCE: &addCarryImmediateInstruction{basicInstruction{8, 1}, regs},
		0x97: &subInstruction{basicInstruction{4, 0}, a, regs},
		0x90: &subInstruction{basicInstruction{4, 0}, b, regs},
		0x91: &subInstruction{basicInstruction{4, 0}, c, regs},
		0x92: &subInstruction{basicInstruction{4, 0}, d, regs},
		0x93: &subInstruction{basicInstruction{4, 0}, e, regs},
		0x94: &subInstruction{basicInstruction{4, 0}, h, regs},
		0x95: &subInstruction{basicInstruction{4, 0}, l, regs},
		0x96: &subFromMemoryInstruction{basicInstruction{8, 0}, h, l, regs, mmu},
		// ??? 0xD6: &subInstruction{8, 1, a, regs},
		0x9F: &subCarryInstruction{basicInstruction{4, 0}, a, regs},
		0x98: &subCarryInstruction{basicInstruction{4, 0}, b, regs},
		0x99: &subCarryInstruction{basicInstruction{4, 0}, c, regs},
		0x9A: &subCarryInstruction{basicInstruction{4, 0}, d, regs},
		0x9B: &subCarryInstruction{basicInstruction{4, 0}, e, regs},
		0x9C: &subCarryInstruction{basicInstruction{4, 0}, h, regs},
		0x9D: &subCarryInstruction{basicInstruction{4, 0}, l, regs},
		0x9E: &subCarryFromMemoryInstruction{basicInstruction{8, 0}, h, l, regs, mmu},
		0xA7: &andInstruction{basicInstruction{4, 0}, a, regs},
		0xA0: &andInstruction{basicInstruction{4, 0}, b, regs},
		0xA1: &andInstruction{basicInstruction{4, 0}, c, regs},
		0xA2: &andInstruction{basicInstruction{4, 0}, d, regs},
		0xA3: &andInstruction{basicInstruction{4, 0}, e, regs},
		0xA4: &andInstruction{basicInstruction{4, 0}, h, regs},
		0xA5: &andInstruction{basicInstruction{4, 0}, l, regs},
		0xA6: &andFromMemoryInstruction{basicInstruction{8, 0}, h, l, regs, mmu},
		// ??? 0xE6: &andInstruction{8, 1, nil??, regs},
		0xB7: &orInstruction{basicInstruction{4, 0}, a, regs},
		0xB0: &orInstruction{basicInstruction{4, 0}, b, regs},
		0xB1: &orInstruction{basicInstruction{4, 0}, c, regs},
		0xB2: &orInstruction{basicInstruction{4, 0}, d, regs},
		0xB3: &orInstruction{basicInstruction{4, 0}, e, regs},
		0xB4: &orInstruction{basicInstruction{4, 0}, h, regs},
		0xB5: &orInstruction{basicInstruction{4, 0}, l, regs},
		0xB6: &orFromMemoryInstruction{basicInstruction{8, 0}, h, l, regs, mmu},
		// ?? 0xF6: &orInstruction{8, 1, nil??, regs},
		0xAF: &xorInstruction{basicInstruction{4, 0}, a, regs},
		0xA8: &xorInstruction{basicInstruction{4, 0}, b, regs},
		0xA9: &xorInstruction{basicInstruction{4, 0}, c, regs},
		0xAA: &xorInstruction{basicInstruction{4, 0}, d, regs},
		0xAB: &xorInstruction{basicInstruction{4, 0}, e, regs},
		0xAC: &xorInstruction{basicInstruction{4, 0}, h, regs},
		0xAD: &xorInstruction{basicInstruction{4, 0}, l, regs},
		// ?? 0xAE: &xorInstruction{8, 0, hl, regs},
		// ?? 0xEE: &xorInstruction{8, 1, nil??, regs},
		0xBF: &cpInstruction{basicInstruction{4, 0}, a, regs},
		0xB8: &cpInstruction{basicInstruction{4, 0}, b, regs},
		0xB9: &cpInstruction{basicInstruction{4, 0}, c, regs},
		0xBA: &cpInstruction{basicInstruction{4, 0}, d, regs},
		0xBB: &cpInstruction{basicInstruction{4, 0}, e, regs},
		0xBC: &cpInstruction{basicInstruction{4, 0}, h, regs},
		0xBD: &cpInstruction{basicInstruction{4, 0}, l, regs},
		// ?? 0xBE: &cpInstruction{8, 0, hl, regs},
		// ?? 0xFE: &cpInstruction{8, 1, nil??, regs},
		0x3C: &incInstruction{basicInstruction{4, 0}, a, regs},
		0x04: &incInstruction{basicInstruction{4, 0}, b, regs},
		0x0C: &incInstruction{basicInstruction{4, 0}, c, regs},
		0x14: &incInstruction{basicInstruction{4, 0}, d, regs},
		0x1C: &incInstruction{basicInstruction{4, 0}, e, regs},
		0x24: &incInstruction{basicInstruction{4, 0}, h, regs},
		0x2C: &incInstruction{basicInstruction{4, 0}, l, regs},
		// ?? 0x34: &incInstruction{12, 0, hl, regs},
		0x3D: &decInstruction{basicInstruction{4, 0}, a, regs},
		0x05: &decInstruction{basicInstruction{4, 0}, b, regs},
		0x0D: &decInstruction{basicInstruction{4, 0}, c, regs},
		0x15: &decInstruction{basicInstruction{4, 0}, d, regs},
		0x1D: &decInstruction{basicInstruction{4, 0}, e, regs},
		0x25: &decInstruction{basicInstruction{4, 0}, h, regs},
		0x2D: &decInstruction{basicInstruction{4, 0}, l, regs},
		// ?? 0x35: &decInstruction{12, 0, hl, regs},
		0x09: &add16BitInstruction{basicInstruction{8, 0}, b, c, regs},
		0x19: &add16BitInstruction{basicInstruction{8, 0}, d, e, regs},
		0x29: &add16BitInstruction{basicInstruction{8, 0}, h, l, regs},
		0x39: &add16BitFromSPInstruction{basicInstruction{8, 0}, regs},
		0xE8: &addSP16BitInstruction{basicInstruction{16, 0}, regs},
		0x03: &inc16BitInstruction{basicInstruction{8, 0}, b, c, regs},
		0x13: &inc16BitInstruction{basicInstruction{8, 0}, d, e, regs},
		0x23: &inc16BitInstruction{basicInstruction{8, 0}, h, l, regs},
		0x33: &incSP16BitInstruction{basicInstruction{8, 0}, regs},
		0x0B: &dec16BitInstruction{basicInstruction{8, 0}, b, c, regs},
		0x1B: &dec16BitInstruction{basicInstruction{8, 0}, d, e, regs},
		0x2B: &dec16BitInstruction{basicInstruction{8, 0}, h, l, regs},
		0x3B: &decSP16BitInstruction{basicInstruction{8, 0}, regs},
		// ?? 0x37: &swapInstruction{8, 0, a, regs},
		0x27: &daaInstruction{basicInstruction{4, 0}, regs},
		0x2F: &cplInstruction{basicInstruction{4, 0}, regs},
		// ?? 0x3F: &ccfInstruction{4, 0 regs},
		// ?? 0x37: &scfInstruction{4, 0, regs},
		0x76: &haltInstruction{basicInstruction{4, 0}},
		// ?? 0x1000: &stopInstruction{4, 0},
		0xF3: &diInstruction{basicInstruction{4, 0}},
		0xFB: &eiInstruction{basicInstruction{4, 0}},
		0xC3: &jumpInstruction{basicInstruction{12, 2}},
		0xC2: &conditionalJumpInstruction{basicInstruction{12, 2}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 0 }},
		0xCA: &conditionalJumpInstruction{basicInstruction{12, 2}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 1 }},
		0xD2: &conditionalJumpInstruction{basicInstruction{12, 2}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 0 }},
		0xDA: &conditionalJumpInstruction{basicInstruction{12, 2}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 1 }},
		0xE9: &jumpHlInstruction{basicInstruction{4, 0}, regs},
		0x18: &jumpImmediateInstruction{basicInstruction{8, 1}, regs},
		0x20: &conditionalJumpImmediateInstruction{basicInstruction{8, 1}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 0 }},
		0x28: &conditionalJumpImmediateInstruction{basicInstruction{8, 1}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 1 }},
		0x30: &conditionalJumpImmediateInstruction{basicInstruction{8, 1}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 0 }},
		0x38: &conditionalJumpImmediateInstruction{basicInstruction{8, 1}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 1 }},
		0xCD: &callInstruction{basicInstruction{12, 2}, regs},
		0xC4: &callConditionalInstruction{basicInstruction{12, 2}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 0 }},
		0xCC: &callConditionalInstruction{basicInstruction{12, 2}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 1 }},
		0xD4: &callConditionalInstruction{basicInstruction{12, 2}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 0 }},
		0xDC: &callConditionalInstruction{basicInstruction{12, 2}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 1 }},
		0xC7: &restartInstruction{basicInstruction{32, 0}, regs, 0x00},
		0xCF: &restartInstruction{basicInstruction{32, 0}, regs, 0x08},
		0xD7: &restartInstruction{basicInstruction{32, 0}, regs, 0x10},
		0xDF: &restartInstruction{basicInstruction{32, 0}, regs, 0x18},
		0xE7: &restartInstruction{basicInstruction{32, 0}, regs, 0x20},
		0xEF: &restartInstruction{basicInstruction{32, 0}, regs, 0x28},
		0xF7: &restartInstruction{basicInstruction{32, 0}, regs, 0x30},
		0xFF: &restartInstruction{basicInstruction{32, 0}, regs, 0x38},
		0xC9: &returnInstruction{basicInstruction{8, 0}, regs},
		0xC0: &returnConditionalInstruction{basicInstruction{8, 0}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 0 }},
		0xC8: &returnConditionalInstruction{basicInstruction{8, 0}, regs, func() bool { return (regs.ReadRegister(f) >> 7) == 1 }},
		0xD0: &returnConditionalInstruction{basicInstruction{8, 0}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 0 }},
		0xD8: &returnConditionalInstruction{basicInstruction{8, 0}, regs, func() bool { return ((regs.ReadRegister(f) & 0x08) >> 3) == 1 }},
		0xD9: &retiInstruction{basicInstruction{8, 0}, regs},
	}
}

// Assuming for now that loading of 2 byte immediate values is handled in different type

func (a *address) ShouldJump() bool {
	return a.pcShouldJump
}

func (a *address) NewAddress() uint16 {
	return a.newAddress
}

func (i *basicInstruction) GetNumParameterBytes() int {
	return i.paramBytes
}

func (n *noopInstruction) Execute(_ Parameters) Addresser {
	return &address{}
}

func (i *loadImmediateInstruction) Execute(params Parameters) Addresser {
	i.regs.WriteRegister(i.dest, params[0])
	return &address{}
}

func (i *loadRegisterInstruction) Execute(params Parameters) Addresser {
	i.regs.WriteRegister(i.dest, i.regs.ReadRegister(i.source))
	return &address{}
}

func (i *loadRegisterWithOffsetInstruction) Execute(params Parameters) Addresser {
	i.regs.WriteRegister(i.dest, i.mmu.ReadAt(i.memAddr+uint16(i.regs.ReadRegister(i.offsetReg))))
	return &address{}
}

func (i *loadMemoryWithRegisterInstruction) Execute(params Parameters) Addresser {
	i.mmu.WriteByte(i.memAddr+uint16(i.regs.ReadRegister(i.offset)), i.regs.ReadRegister(i.source))
	return &address{}
}

func (i *loadTwoByteImmediateInstruction) Execute(params Parameters) Addresser {
	i.regs.WriteRegister(i.dest1, params[0])
	i.regs.WriteRegister(i.dest2, params[1])
	return &address{}
}

func (i *ldsphlInstruction) Execute(params Parameters) Addresser {
	val, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	i.regs.WriteSP(val)
	return &address{}
}

func (i *ldhImmediateInstruction) Execute(params Parameters) Addresser {
	i.mmu.WriteByte(0xFF00+uint16(params[0]), i.regs.ReadRegister(i.source))
	return &address{}
}

func (i *ldhaInstruction) Execute(params Parameters) Addresser {
	val := i.mmu.ReadAt(0xFF00 + uint16(params[0]))
	i.regs.WriteRegister(i.dest, val)
	return &address{}
}

func (i *ldiInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	i.regs.WriteRegister(i.dest, val)
	i.regs.WriteRegisterPair(i.source1, i.source2, addr+1)
	return &address{}
}

func (i *ldihlInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(a)
	addr, err := i.regs.ReadRegisterPair(i.dest1, i.dest2)
	if err != nil {
		fmt.Println(err)
	}
	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegisterPair(i.dest1, i.dest2, addr+1)
	return &address{}
}

func (i *lddhlInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(a)
	addr, err := i.regs.ReadRegisterPair(i.dest1, i.dest2)
	if err != nil {
		fmt.Println(err)
	}
	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegisterPair(i.dest1, i.dest2, addr-1)
	return &address{}
}

func (i *lddInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	i.regs.WriteRegister(i.dest, val)
	i.regs.WriteRegisterPair(i.source1, i.source2, addr-1)
	return &address{}
}

func (i *loadRegisterFromMemoryInstruction) Execute(params Parameters) Addresser {
	val, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	memVal := i.mmu.ReadAt(val)
	i.regs.WriteRegister(i.dest, memVal)
	return &address{}
}

func (i *pushInstruction) Execute(params Parameters) Addresser {
	val, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	i.regs.PushSP(val)
	return &address{}
}

func (i *writeMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.dest1, i.dest2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.regs.ReadRegister(i.source)
	i.mmu.WriteByte(addr, val)
	return &address{}
}

func (i *popInstruction) Execute(params Parameters) Addresser {
	stackValue := i.regs.PopSP()
	i.regs.WriteRegister(i.source1, byte((stackValue&0xFF00)>>8))
	i.regs.WriteRegister(i.source2, byte(stackValue&0x00FF))
	return &address{}
}

func (i *addImmediateInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	aValue := i.regs.ReadRegister(a)
	result := aValue + byte(params[0])
	if result == 0 {
		flags += 0x80
	}
	if flags&0x0F == 0 && params[0] != 0 {
		flags += 0x20
	}
	if flags&0xFF == 0 && params[0] != 0 {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *addCarryImmediateInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	carryBit := ((flags & 0x10) >> 4)
	aValue := i.regs.ReadRegister(a)
	result := aValue + byte(params[0]) + carryBit
	if result == 0 {
		flags += 0x80
	}
	if flags&0x0F == 0 && params[0] != 0 {
		flags += 0x20
	}
	if flags&0xFF == 0 && params[0] != 0 {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *addCarryFromMemoryInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	carryBit := ((flags & 0x10) >> 4)
	aValue := i.regs.ReadRegister(i.dest)
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	result := aValue + val + carryBit
	if result == 0 {
		flags += 0x80
	}
	if flags&0x0F == 0 && val != 0 {
		flags += 0x20
	}
	if flags&0xFF == 0 && val != 0 {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *subCarryFromMemoryInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	carryBit := ((flags & 0x10) >> 4)
	aValue := i.regs.ReadRegister(a)
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	subtractor := i.mmu.ReadAt(addr)
	result := aValue - subtractor - carryBit
	flags = 0x40
	if result == 0 {
		flags += 0x80
	}
	if result&0x0F == 0x0F {
		flags += 0x20
	}
	if result&0xFF == 0xFF {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *subFromMemoryInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	aValue := i.regs.ReadRegister(a)
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	subtractor := i.mmu.ReadAt(addr)
	result := aValue - subtractor
	flags = 0x40
	if result == 0 {
		flags += 0x80
	}
	if result&0x0F == 0x0F {
		flags += 0x20
	}
	if result&0xFF == 0xFF {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *addFromMemoryInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	aValue := i.regs.ReadRegister(a)
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	result := aValue + val
	if result == 0 {
		flags += 0x80
	}
	if flags&0x0F == 0 && val != 0 {
		flags += 0x20
	}
	if flags&0xFF == 0 && val != 0 {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *addInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	adder := i.regs.ReadRegister(i.source)
	aValue := i.regs.ReadRegister(a)
	result := aValue + adder
	flags = 0
	if result == 0 {
		flags += 0x80
	}
	if result&0x0F == 0 && adder != 0 {
		flags += 0x20
	}
	if result&0xFF == 0 && adder != 0 {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *addCarryInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	adder := i.regs.ReadRegister(i.source)
	aValue := i.regs.ReadRegister(a)
	carryBit := ((flags & 0x10) >> 4)
	result := aValue + adder + carryBit
	flags = 0
	if result == 0 {
		flags += 0x80
	}
	if result&0x0F == 0 && (carryBit != 0 || adder != 0) {
		flags += 0x20
	}
	if result&0xFF == 0 && (carryBit != 0 || adder != 0) {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *subInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	subtractor := i.regs.ReadRegister(i.source)
	aValue := i.regs.ReadRegister(a)
	result := aValue - subtractor
	flags = 0x40
	if result == 0 {
		flags += 0x80
	}
	if result&0x0F == 0x0F {
		flags += 0x20
	}
	if result&0xFF == 0xFF {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *subCarryInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	subtractor := i.regs.ReadRegister(i.source)
	aValue := i.regs.ReadRegister(a)
	result := aValue - subtractor - ((flags & 0x10) >> 4)
	flags = 0x40
	if result == 0 {
		flags += 0x80
	}
	if result&0x0F == 0x0F {
		flags += 0x20
	}
	if result&0xFF == 0xFF {
		flags += 0x10
	}
	i.regs.WriteRegister(a, result)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *andInstruction) Execute(params Parameters) Addresser {
	result := i.regs.ReadRegister(a) & i.regs.ReadRegister(i.source)
	flags := 0
	if result == 0 {
		flags = 0x80
	}
	i.regs.WriteRegister(f, byte(flags))
	i.regs.WriteRegister(a, result)
	return &address{}
}

func (i *andFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	result := i.regs.ReadRegister(a) & val
	flags := 0
	if result == 0 {
		flags = 0x80
	}
	i.regs.WriteRegister(f, byte(flags))
	i.regs.WriteRegister(a, result)
	return &address{}
}

func (i *orFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	result := i.regs.ReadRegister(a) | val
	flags := 0
	if result == 0 {
		flags = 0x80
	}
	i.regs.WriteRegister(f, byte(flags))
	i.regs.WriteRegister(a, result)
	return &address{}
}

func (i *orInstruction) Execute(params Parameters) Addresser {
	result := i.regs.ReadRegister(a) | i.regs.ReadRegister(i.source)
	flags := 0
	if result == 0 {
		flags = 0x80
	}
	i.regs.WriteRegister(f, byte(flags))
	i.regs.WriteRegister(a, result)
	return &address{}
}

func (i *xorInstruction) Execute(params Parameters) Addresser {
	result := i.regs.ReadRegister(a) ^ i.regs.ReadRegister(i.source)
	flags := 0
	if result == 0 {
		flags = 0x80
	}
	i.regs.WriteRegister(f, byte(flags))
	i.regs.WriteRegister(a, result)
	return &address{}
}

func (i *cpInstruction) Execute(params Parameters) Addresser {
	aValue := i.regs.ReadRegister(a)
	otherValue := i.regs.ReadRegister(i.source)
	result := aValue - otherValue
	flags := i.regs.ReadRegister(f)
	flags = flags | 0x40
	if result == 0 {
		flags += 0x80
	}
	if result&0x0F == 0x0F {
		flags += 0x20
	}
	if aValue < otherValue {
		flags += 0x10
	}
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *incInstruction) Execute(params Parameters) Addresser {
	newValue := i.regs.ReadRegister(i.source) + 1
	flags := i.regs.ReadRegister(f)
	flags = flags & 0x10
	if newValue == 0 {
		flags += 0x80
	}
	if newValue&0x0F == 0 {
		flags += 0x20
	}
	i.regs.WriteRegister(i.source, newValue)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *decInstruction) Execute(params Parameters) Addresser {
	newValue := i.regs.ReadRegister(i.source) - 1
	flags := i.regs.ReadRegister(f)
	flags = flags | 0x40
	if newValue == 0 {
		flags += 0x80
	}
	if newValue&0x0F == 0x0F {
		flags += 0x20
	}
	i.regs.WriteRegister(i.source, newValue)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

// TODO: test this instruction with the debugger
func (i *ldhlspInstruction) Execute(params Parameters) Addresser {
	immediateValue := computeTwosComplement(uint8(params[0]))
	spValue := int(i.regs.PopSP())
	flags := 0
	if immediateValue < 0 {
		if ((spValue + immediateValue) & 0xFF) > (spValue & 0xFF) {
			flags += 1
		}
	} else {
		if ((spValue + immediateValue) & 0xFF) < (spValue & 0xFF) {
			flags += 1 // set c flag
		}
		if ((spValue + immediateValue) & 0x0F) < (spValue & 0x0F) {
			flags += 2 // set h flag
		}
	}
	spValue += immediateValue
	i.regs.WriteRegisterPair(h, l, uint16(spValue)+uint16(immediateValue))
	i.regs.WriteRegister(f, uint8(flags))
	return &address{}
}

func (i *addSP16BitInstruction) Execute(params Parameters) Addresser {
	immediateValue := computeTwosComplement(uint8(params[0]))
	spValue := int(i.regs.PopSP())
	flags := 0

	if immediateValue < 0 {
		if ((spValue + immediateValue) & 0xFF) > (spValue & 0xFF) {
			flags += 1 // set c flag
		}
		// TODO: based on behavior from the debugger I am currently using it doesn't
		// look like this needs to be set when subtracting????
		//	if ((spValue + immediateValue) & 0x0F) > (spValue & 0x0F) {
		//		flags += 2 // set h flag
		//	}

	} else {
		if ((spValue + immediateValue) & 0xFF) < (spValue & 0xFF) {
			flags += 1 // set c flag
		}
		if ((spValue + immediateValue) & 0x0F) < (spValue & 0x0F) {
			flags += 2 // set h flag
		}

	}
	spValue += immediateValue
	i.regs.WriteSP(uint16(spValue))
	i.regs.WriteRegister(f, uint8(flags))
	return &address{}
}

func (i *inc16BitInstruction) Execute(params Parameters) Addresser {
	val, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val += 1
	i.regs.WriteRegisterPair(i.source1, i.source2, val)
	return &address{}
}

func (i *incSP16BitInstruction) Execute(params Parameters) Addresser {
	i.regs.WriteSP(i.regs.PopSP() + 1)
	return &address{}
}

func (i *add16BitFromSPInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	flags = flags & 0x80

	spVal := i.regs.PopSP()
	hlVal, err := i.regs.ReadRegisterPair(h, l)
	if err != nil {
		fmt.Println(err)
	}

	if spVal+hlVal < hlVal {
		flags += 3
	}
	i.regs.WriteRegisterPair(h, l, hlVal+spVal)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *add16BitInstruction) Execute(params Parameters) Addresser {
	flags := i.regs.ReadRegister(f)
	flags = flags & 0x80
	val, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	hlVal, err := i.regs.ReadRegisterPair(h, l)
	if err != nil {
		fmt.Println(err)
	}

	if val+hlVal < hlVal {
		flags += 3
	}
	i.regs.WriteRegisterPair(h, l, val+hlVal)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *decSP16BitInstruction) Execute(params Parameters) Addresser {
	i.regs.WriteSP(i.regs.PopSP() - 1)
	return &address{}
}

func (i *dec16BitInstruction) Execute(params Parameters) Addresser {
	val, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val -= 1
	i.regs.WriteRegisterPair(i.source1, i.source2, uint16(val))
	return &address{}
}

func (i *daaInstruction) Execute(params Parameters) Addresser {
	aValue := i.regs.ReadRegister(a)
	flagValue := i.regs.ReadRegister(f)
	flags := flagValue & 0x40 // Preserve n flag but pre-emptively reset others
	carryPerformed := false   // Pretty weird but the A register knows if it overflows above 0xFF and will
	// continue the daa with that knowledge. Since golang will overflow our byte
	// value we need to know whether a 0 value was caused by carrying over a digit
	// from the previous operation or whether it was 0 to begin with

	if aValue&0x0F > 9 || (flagValue&0x20) == 0x20 {
		aValue = aValue + 0x06
		carryPerformed = true
	}

	if ((aValue&0xF0)>>4) > 9 || (aValue == 0 && carryPerformed) || (flagValue&0x10) == 0x10 {
		if aValue+0x60 < aValue { // Check for overflow of upper 4 bits
			flags += 1
		}
		aValue = aValue + 0x60
	}

	if aValue == 0 {
		flags += 8
	}
	i.regs.WriteRegister(a, aValue)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

// A & 0xFF, also set n, h flags to 1
func (i *cplInstruction) Execute(params Parameters) Addresser {
	i.regs.WriteRegister(a, (i.regs.ReadRegister(a) ^ 0xFF))
	i.regs.WriteRegister(f, (i.regs.ReadRegister(f) | 0x60))
	return &address{}
}

func (i *haltInstruction) Execute(params Parameters) Addresser {
	return &address{}
}

func (i *stopInstruction) Execute(params Parameters) Addresser {
	return &address{}
}

func (i *diInstruction) Execute(params Parameters) Addresser {
	return &address{}
}

func (i *eiInstruction) Execute(params Parameters) Addresser {
	return &address{}
}

func (i *jumpInstruction) Execute(params Parameters) Addresser {
	newPC := (uint16(params[1]) << 0xFF) + uint16(params[0])
	return &address{true, newPC}
}

func (i *conditionalJumpInstruction) Execute(params Parameters) Addresser {
	if i.conditional() == true {
		newPC := (uint16(params[1]) << 0xFF) + uint16(params[0])
		return &address{true, newPC}
	}
	return &address{}
}

func (i *jumpHlInstruction) Execute(params Parameters) Addresser {
	newPC, err := i.regs.ReadRegisterPair(h, l)
	if err != nil {
		fmt.Println("Error occured when attempting to execute jump HL instruction: ", err)
		return &address{}
	}
	return &address{true, newPC}
}

func (i *jumpImmediateInstruction) Execute(params Parameters) Addresser {
	newPC := i.regs.ReadPC() + uint16(params[0])
	return &address{true, newPC}
}

func (i *conditionalJumpImmediateInstruction) Execute(params Parameters) Addresser {
	if i.conditional() == true {
		newPC := i.regs.ReadPC() + uint16(params[0])
		return &address{true, newPC}
	}
	return &address{}
}

func (i *callInstruction) Execute(params Parameters) Addresser {
	i.regs.PushSP(i.regs.ReadPC()) // TODO: This may need to increment PC before saving
	newPC := (uint16(params[1]) << 0xFF) + uint16(params[0])
	return &address{true, newPC}
}

func (i *callConditionalInstruction) Execute(params Parameters) Addresser {
	if i.conditional() == true {
		i.regs.PushSP(i.regs.ReadPC()) // TODO: This may need to increment PC before saving
		newPC := (uint16(params[1]) << 0xFF) + uint16(params[0])
		return &address{true, newPC}
	}
	return &address{}
}

func (i *restartInstruction) Execute(params Parameters) Addresser {
	i.regs.PushSP(i.regs.ReadPC()) // TODO: This may need to increment PC before saving
	return &address{true, 0x00 + i.offset}
}

func (i *returnInstruction) Execute(params Parameters) Addresser {
	newPC := i.regs.PopSP()
	return &address{true, newPC}
}

func (i *returnConditionalInstruction) Execute(params Parameters) Addresser {
	if i.conditional() == true {
		newPC := i.regs.PopSP()
		return &address{true, newPC}
	}
	return &address{}
}

func (i *retiInstruction) Execute(params Parameters) Addresser {
	// Need to deal with enabling/disabling interrupts
	newPC := i.regs.PopSP()
	return &address{true, newPC}
}

func computeTwosComplement(number uint8) int {
	mask := (1 << 7)
	return int(number & ^uint8(mask)) - int(uint8(mask)&number)
}
