package main

import "fmt"

type ExtendedInstruction interface {
	Execute(Parameters) Addresser
	GetNumParameterBytes() int
	GetCycles(Parameters) int
}

type swapInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type swapFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type rlcInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type rlcFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type rlInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type rlFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type rrcInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type rrcFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type rrInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type rrFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type slaInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type slaFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type sraInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type sraFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type srlInstruction struct {
	basicInstruction
	source Register
	regs   Registers
}

type srlFromMemoryInstruction struct {
	basicInstruction
	source1 Register
	source2 Register
	regs    Registers
	mmu     MMU
}

type bitInstruction struct {
	basicInstruction
	bitNumber uint
	source    Register
	regs      Registers
}

type bitFromMemoryInstruction struct {
	basicInstruction
	bitNumber uint
	source1   Register
	source2   Register
	regs      Registers
	mmu       MMU
}

type setInstruction struct {
	basicInstruction
	bitNumber uint
	source    Register
	regs      Registers
}

type setFromMemoryInstruction struct {
	basicInstruction
	bitNumber uint
	source1   Register
	source2   Register
	regs      Registers
	mmu       MMU
}

type resetInstruction struct {
	basicInstruction
	bitNumber uint
	source    Register
	regs      Registers
}

type resetFromMemoryInstruction struct {
	basicInstruction
	bitNumber uint
	source1   Register
	source2   Register
	regs      Registers
	mmu       MMU
}

func CreateExtendedInstructions(regs Registers, mmu MMU) map[byte]ExtendedInstruction {
	return map[byte]ExtendedInstruction{
		0x00: &rlcInstruction{basicInstruction{8, 0}, b, regs},
		0x01: &rlcInstruction{basicInstruction{8, 0}, c, regs},
		0x02: &rlcInstruction{basicInstruction{8, 0}, d, regs},
		0x03: &rlcInstruction{basicInstruction{8, 0}, e, regs},
		0x04: &rlcInstruction{basicInstruction{8, 0}, h, regs},
		0x05: &rlcInstruction{basicInstruction{8, 0}, l, regs},
		0x06: &rlcFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x07: &rlcInstruction{basicInstruction{8, 0}, a, regs},
		0x10: &rlInstruction{basicInstruction{8, 0}, b, regs},
		0x11: &rlInstruction{basicInstruction{8, 0}, c, regs},
		0x12: &rlInstruction{basicInstruction{8, 0}, d, regs},
		0x13: &rlInstruction{basicInstruction{8, 0}, e, regs},
		0x14: &rlInstruction{basicInstruction{8, 0}, h, regs},
		0x15: &rlInstruction{basicInstruction{8, 0}, l, regs},
		0x16: &rlFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x17: &rlInstruction{basicInstruction{8, 0}, a, regs},
		0x08: &rrcInstruction{basicInstruction{8, 0}, b, regs},
		0x09: &rrcInstruction{basicInstruction{8, 0}, c, regs},
		0x0A: &rrcInstruction{basicInstruction{8, 0}, d, regs},
		0x0B: &rrcInstruction{basicInstruction{8, 0}, e, regs},
		0x0C: &rrcInstruction{basicInstruction{8, 0}, h, regs},
		0x0D: &rrcInstruction{basicInstruction{8, 0}, l, regs},
		0x0E: &rrcFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x0F: &rrcInstruction{basicInstruction{8, 0}, a, regs},
		0x18: &rrInstruction{basicInstruction{8, 0}, b, regs},
		0x19: &rrInstruction{basicInstruction{8, 0}, c, regs},
		0x1A: &rrInstruction{basicInstruction{8, 0}, d, regs},
		0x1B: &rrInstruction{basicInstruction{8, 0}, e, regs},
		0x1C: &rrInstruction{basicInstruction{8, 0}, h, regs},
		0x1D: &rrInstruction{basicInstruction{8, 0}, l, regs},
		0x1E: &rrFromMemoryInstruction{basicInstruction{8, 0}, h, l, regs, mmu},
		0x1F: &rrInstruction{basicInstruction{8, 0}, a, regs},
		0x20: &slaInstruction{basicInstruction{8, 0}, b, regs},
		0x21: &slaInstruction{basicInstruction{8, 0}, c, regs},
		0x22: &slaInstruction{basicInstruction{8, 0}, d, regs},
		0x23: &slaInstruction{basicInstruction{8, 0}, e, regs},
		0x24: &slaInstruction{basicInstruction{8, 0}, h, regs},
		0x25: &slaInstruction{basicInstruction{8, 0}, l, regs},
		0x26: &slaFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x27: &slaInstruction{basicInstruction{8, 0}, a, regs},
		0x28: &sraInstruction{basicInstruction{8, 0}, b, regs},
		0x29: &sraInstruction{basicInstruction{8, 0}, c, regs},
		0x2A: &sraInstruction{basicInstruction{8, 0}, d, regs},
		0x2B: &sraInstruction{basicInstruction{8, 0}, e, regs},
		0x2C: &sraInstruction{basicInstruction{8, 0}, h, regs},
		0x2D: &sraInstruction{basicInstruction{8, 0}, l, regs},
		0x2E: &sraFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x2F: &sraInstruction{basicInstruction{8, 0}, a, regs},
		0x30: &swapInstruction{basicInstruction{8, 0}, b, regs},
		0x31: &swapInstruction{basicInstruction{8, 0}, c, regs},
		0x32: &swapInstruction{basicInstruction{8, 0}, d, regs},
		0x33: &swapInstruction{basicInstruction{8, 0}, e, regs},
		0x34: &swapInstruction{basicInstruction{8, 0}, h, regs},
		0x35: &swapInstruction{basicInstruction{8, 0}, l, regs},
		0x36: &swapFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x37: &swapInstruction{basicInstruction{8, 0}, a, regs},
		0x38: &srlInstruction{basicInstruction{8, 0}, b, regs},
		0x39: &srlInstruction{basicInstruction{8, 0}, c, regs},
		0x3A: &srlInstruction{basicInstruction{8, 0}, d, regs},
		0x3B: &srlInstruction{basicInstruction{8, 0}, e, regs},
		0x3C: &srlInstruction{basicInstruction{8, 0}, h, regs},
		0x3D: &srlInstruction{basicInstruction{8, 0}, l, regs},
		0x3E: &srlFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x3F: &srlInstruction{basicInstruction{8, 0}, a, regs},
		0x40: &bitInstruction{basicInstruction{8, 0}, 0, b, regs},
		0x41: &bitInstruction{basicInstruction{8, 0}, 0, c, regs},
		0x42: &bitInstruction{basicInstruction{8, 0}, 0, d, regs},
		0x43: &bitInstruction{basicInstruction{8, 0}, 0, e, regs},
		0x44: &bitInstruction{basicInstruction{8, 0}, 0, h, regs},
		0x45: &bitInstruction{basicInstruction{8, 0}, 0, l, regs},
		0x46: &bitFromMemoryInstruction{basicInstruction{16, 0}, 0, h, l, regs, mmu},
		0x47: &bitInstruction{basicInstruction{8, 0}, 0, a, regs},
		0x48: &bitInstruction{basicInstruction{8, 0}, 1, b, regs},
		0x49: &bitInstruction{basicInstruction{8, 0}, 1, c, regs},
		0x4A: &bitInstruction{basicInstruction{8, 0}, 1, d, regs},
		0x4B: &bitInstruction{basicInstruction{8, 0}, 1, e, regs},
		0x4C: &bitInstruction{basicInstruction{8, 0}, 1, h, regs},
		0x4D: &bitInstruction{basicInstruction{8, 0}, 1, l, regs},
		0x4E: &bitFromMemoryInstruction{basicInstruction{16, 0}, 1, h, l, regs, mmu},
		0x4F: &bitInstruction{basicInstruction{8, 0}, 1, a, regs},
		0x50: &bitInstruction{basicInstruction{8, 0}, 2, b, regs},
		0x51: &bitInstruction{basicInstruction{8, 0}, 2, c, regs},
		0x52: &bitInstruction{basicInstruction{8, 0}, 2, d, regs},
		0x53: &bitInstruction{basicInstruction{8, 0}, 2, e, regs},
		0x54: &bitInstruction{basicInstruction{8, 0}, 2, h, regs},
		0x55: &bitInstruction{basicInstruction{8, 0}, 2, l, regs},
		0x56: &bitFromMemoryInstruction{basicInstruction{16, 0}, 2, h, l, regs, mmu},
		0x57: &bitInstruction{basicInstruction{8, 0}, 2, a, regs},
		0x58: &bitInstruction{basicInstruction{8, 0}, 3, b, regs},
		0x59: &bitInstruction{basicInstruction{8, 0}, 3, c, regs},
		0x5A: &bitInstruction{basicInstruction{8, 0}, 3, d, regs},
		0x5B: &bitInstruction{basicInstruction{8, 0}, 3, e, regs},
		0x5C: &bitInstruction{basicInstruction{8, 0}, 3, h, regs},
		0x5D: &bitInstruction{basicInstruction{8, 0}, 3, l, regs},
		0x5E: &bitFromMemoryInstruction{basicInstruction{16, 0}, 3, h, l, regs, mmu},
		0x5F: &bitInstruction{basicInstruction{8, 0}, 3, a, regs},
		0x60: &bitInstruction{basicInstruction{8, 0}, 4, b, regs},
		0x61: &bitInstruction{basicInstruction{8, 0}, 4, c, regs},
		0x62: &bitInstruction{basicInstruction{8, 0}, 4, d, regs},
		0x63: &bitInstruction{basicInstruction{8, 0}, 4, e, regs},
		0x64: &bitInstruction{basicInstruction{8, 0}, 4, h, regs},
		0x65: &bitInstruction{basicInstruction{8, 0}, 4, l, regs},
		0x66: &bitFromMemoryInstruction{basicInstruction{16, 0}, 4, h, l, regs, mmu},
		0x67: &bitInstruction{basicInstruction{8, 0}, 4, a, regs},
		0x68: &bitInstruction{basicInstruction{8, 0}, 5, b, regs},
		0x69: &bitInstruction{basicInstruction{8, 0}, 5, c, regs},
		0x6A: &bitInstruction{basicInstruction{8, 0}, 5, d, regs},
		0x6B: &bitInstruction{basicInstruction{8, 0}, 5, e, regs},
		0x6C: &bitInstruction{basicInstruction{8, 0}, 5, h, regs},
		0x6D: &bitInstruction{basicInstruction{8, 0}, 5, l, regs},
		0x6E: &bitFromMemoryInstruction{basicInstruction{16, 0}, 5, h, l, regs, mmu},
		0x6F: &bitInstruction{basicInstruction{8, 0}, 5, a, regs},
		0x70: &bitInstruction{basicInstruction{8, 0}, 6, b, regs},
		0x71: &bitInstruction{basicInstruction{8, 0}, 6, c, regs},
		0x72: &bitInstruction{basicInstruction{8, 0}, 6, d, regs},
		0x73: &bitInstruction{basicInstruction{8, 0}, 6, e, regs},
		0x74: &bitInstruction{basicInstruction{8, 0}, 6, h, regs},
		0x75: &bitInstruction{basicInstruction{8, 0}, 6, l, regs},
		0x76: &bitFromMemoryInstruction{basicInstruction{16, 0}, 6, h, l, regs, mmu},
		0x77: &bitInstruction{basicInstruction{8, 0}, 6, a, regs},
		0x78: &bitInstruction{basicInstruction{8, 0}, 7, b, regs},
		0x79: &bitInstruction{basicInstruction{8, 0}, 7, c, regs},
		0x7A: &bitInstruction{basicInstruction{8, 0}, 7, d, regs},
		0x7B: &bitInstruction{basicInstruction{8, 0}, 7, e, regs},
		0x7C: &bitInstruction{basicInstruction{8, 0}, 7, h, regs},
		0x7D: &bitInstruction{basicInstruction{8, 0}, 7, l, regs},
		0x7E: &bitFromMemoryInstruction{basicInstruction{16, 0}, 7, h, l, regs, mmu},
		0x7F: &bitInstruction{basicInstruction{8, 0}, 7, a, regs},
		0x80: &resetInstruction{basicInstruction{8, 0}, 0, b, regs},
		0x81: &resetInstruction{basicInstruction{8, 0}, 0, c, regs},
		0x82: &resetInstruction{basicInstruction{8, 0}, 0, d, regs},
		0x83: &resetInstruction{basicInstruction{8, 0}, 0, e, regs},
		0x84: &resetInstruction{basicInstruction{8, 0}, 0, h, regs},
		0x85: &resetInstruction{basicInstruction{8, 0}, 0, l, regs},
		0x86: &resetFromMemoryInstruction{basicInstruction{16, 0}, 0, h, l, regs, mmu},
		0x87: &resetInstruction{basicInstruction{8, 1}, 0, a, regs},
		0x88: &resetInstruction{basicInstruction{8, 0}, 1, b, regs},
		0x89: &resetInstruction{basicInstruction{8, 0}, 1, c, regs},
		0x8A: &resetInstruction{basicInstruction{8, 0}, 1, d, regs},
		0x8B: &resetInstruction{basicInstruction{8, 0}, 1, e, regs},
		0x8C: &resetInstruction{basicInstruction{8, 0}, 1, h, regs},
		0x8D: &resetInstruction{basicInstruction{8, 0}, 1, l, regs},
		0x8E: &resetFromMemoryInstruction{basicInstruction{16, 0}, 1, h, l, regs, mmu},
		0x8F: &resetInstruction{basicInstruction{8, 0}, 1, a, regs},
		0x90: &resetInstruction{basicInstruction{8, 0}, 2, b, regs},
		0x91: &resetInstruction{basicInstruction{8, 0}, 2, c, regs},
		0x92: &resetInstruction{basicInstruction{8, 0}, 2, d, regs},
		0x93: &resetInstruction{basicInstruction{8, 0}, 2, e, regs},
		0x94: &resetInstruction{basicInstruction{8, 0}, 2, h, regs},
		0x95: &resetInstruction{basicInstruction{8, 0}, 2, l, regs},
		0x96: &resetFromMemoryInstruction{basicInstruction{16, 0}, 2, h, l, regs, mmu},
		0x97: &resetInstruction{basicInstruction{8, 1}, 2, a, regs},
		0x98: &resetInstruction{basicInstruction{8, 0}, 3, b, regs},
		0x99: &resetInstruction{basicInstruction{8, 0}, 3, c, regs},
		0x9A: &resetInstruction{basicInstruction{8, 0}, 3, d, regs},
		0x9B: &resetInstruction{basicInstruction{8, 0}, 3, e, regs},
		0x9C: &resetInstruction{basicInstruction{8, 0}, 3, h, regs},
		0x9D: &resetInstruction{basicInstruction{8, 0}, 3, l, regs},
		0x9E: &resetFromMemoryInstruction{basicInstruction{16, 0}, 3, h, l, regs, mmu},
		0x9F: &resetInstruction{basicInstruction{8, 0}, 3, a, regs},
		0xA0: &resetInstruction{basicInstruction{8, 0}, 4, b, regs},
		0xA1: &resetInstruction{basicInstruction{8, 0}, 4, c, regs},
		0xA2: &resetInstruction{basicInstruction{8, 0}, 4, d, regs},
		0xA3: &resetInstruction{basicInstruction{8, 0}, 4, e, regs},
		0xA4: &resetInstruction{basicInstruction{8, 0}, 4, h, regs},
		0xA5: &resetInstruction{basicInstruction{8, 0}, 4, l, regs},
		0xA6: &resetFromMemoryInstruction{basicInstruction{16, 0}, 4, h, l, regs, mmu},
		0xA7: &resetInstruction{basicInstruction{8, 1}, 4, a, regs},
		0xA8: &resetInstruction{basicInstruction{8, 0}, 5, b, regs},
		0xA9: &resetInstruction{basicInstruction{8, 0}, 5, c, regs},
		0xAA: &resetInstruction{basicInstruction{8, 0}, 5, d, regs},
		0xAB: &resetInstruction{basicInstruction{8, 0}, 5, e, regs},
		0xAC: &resetInstruction{basicInstruction{8, 0}, 5, h, regs},
		0xAD: &resetInstruction{basicInstruction{8, 0}, 5, l, regs},
		0xAE: &resetFromMemoryInstruction{basicInstruction{16, 0}, 5, h, l, regs, mmu},
		0xAF: &resetInstruction{basicInstruction{8, 0}, 5, a, regs},
		0xB0: &resetInstruction{basicInstruction{8, 0}, 6, b, regs},
		0xB1: &resetInstruction{basicInstruction{8, 0}, 6, c, regs},
		0xB2: &resetInstruction{basicInstruction{8, 0}, 6, d, regs},
		0xB3: &resetInstruction{basicInstruction{8, 0}, 6, e, regs},
		0xB4: &resetInstruction{basicInstruction{8, 0}, 6, h, regs},
		0xB5: &resetInstruction{basicInstruction{8, 0}, 6, l, regs},
		0xB6: &resetFromMemoryInstruction{basicInstruction{16, 0}, 6, h, l, regs, mmu},
		0xB7: &resetInstruction{basicInstruction{8, 1}, 6, a, regs},
		0xB8: &resetInstruction{basicInstruction{8, 0}, 7, b, regs},
		0xB9: &resetInstruction{basicInstruction{8, 0}, 7, c, regs},
		0xBA: &resetInstruction{basicInstruction{8, 0}, 7, d, regs},
		0xBB: &resetInstruction{basicInstruction{8, 0}, 7, e, regs},
		0xBC: &resetInstruction{basicInstruction{8, 0}, 7, h, regs},
		0xBD: &resetInstruction{basicInstruction{8, 0}, 7, l, regs},
		0xBE: &resetFromMemoryInstruction{basicInstruction{16, 0}, 7, h, l, regs, mmu},
		0xBF: &resetInstruction{basicInstruction{8, 0}, 7, a, regs},
		0xC0: &setInstruction{basicInstruction{8, 0}, 0, b, regs},
		0xC1: &setInstruction{basicInstruction{8, 0}, 0, c, regs},
		0xC2: &setInstruction{basicInstruction{8, 0}, 0, d, regs},
		0xC3: &setInstruction{basicInstruction{8, 0}, 0, e, regs},
		0xC4: &setInstruction{basicInstruction{8, 0}, 0, h, regs},
		0xC5: &setInstruction{basicInstruction{8, 0}, 0, l, regs},
		0xC6: &setFromMemoryInstruction{basicInstruction{16, 0}, 0, h, l, regs, mmu},
		0xC7: &setInstruction{basicInstruction{8, 0}, 0, a, regs},
		0xC8: &setInstruction{basicInstruction{8, 0}, 1, b, regs},
		0xC9: &setInstruction{basicInstruction{8, 0}, 1, c, regs},
		0xCA: &setInstruction{basicInstruction{8, 0}, 1, d, regs},
		0xCB: &setInstruction{basicInstruction{8, 0}, 1, e, regs},
		0xCC: &setInstruction{basicInstruction{8, 0}, 1, h, regs},
		0xCD: &setInstruction{basicInstruction{8, 0}, 1, l, regs},
		0xCE: &setFromMemoryInstruction{basicInstruction{16, 0}, 1, h, l, regs, mmu},
		0xCF: &setInstruction{basicInstruction{8, 0}, 1, a, regs},
		0xD0: &setInstruction{basicInstruction{8, 0}, 2, b, regs},
		0xD1: &setInstruction{basicInstruction{8, 0}, 2, c, regs},
		0xD2: &setInstruction{basicInstruction{8, 0}, 2, d, regs},
		0xD3: &setInstruction{basicInstruction{8, 0}, 2, e, regs},
		0xD4: &setInstruction{basicInstruction{8, 0}, 2, h, regs},
		0xD5: &setInstruction{basicInstruction{8, 0}, 2, l, regs},
		0xD6: &setFromMemoryInstruction{basicInstruction{16, 0}, 2, h, l, regs, mmu},
		0xD7: &setInstruction{basicInstruction{8, 0}, 2, a, regs},
		0xD8: &setInstruction{basicInstruction{8, 0}, 3, b, regs},
		0xD9: &setInstruction{basicInstruction{8, 0}, 3, c, regs},
		0xDA: &setInstruction{basicInstruction{8, 0}, 3, d, regs},
		0xDB: &setInstruction{basicInstruction{8, 0}, 3, e, regs},
		0xDC: &setInstruction{basicInstruction{8, 0}, 3, h, regs},
		0xDD: &setInstruction{basicInstruction{8, 0}, 3, l, regs},
		0xDE: &setFromMemoryInstruction{basicInstruction{16, 0}, 3, h, l, regs, mmu},
		0xDF: &setInstruction{basicInstruction{8, 0}, 3, a, regs},
		0xE0: &setInstruction{basicInstruction{8, 0}, 4, b, regs},
		0xE1: &setInstruction{basicInstruction{8, 0}, 4, c, regs},
		0xE2: &setInstruction{basicInstruction{8, 0}, 4, d, regs},
		0xE3: &setInstruction{basicInstruction{8, 0}, 4, e, regs},
		0xE4: &setInstruction{basicInstruction{8, 0}, 4, h, regs},
		0xE5: &setInstruction{basicInstruction{8, 0}, 4, l, regs},
		0xE6: &setFromMemoryInstruction{basicInstruction{16, 0}, 4, h, l, regs, mmu},
		0xE7: &setInstruction{basicInstruction{8, 0}, 4, a, regs},
		0xE8: &setInstruction{basicInstruction{8, 0}, 5, b, regs},
		0xE9: &setInstruction{basicInstruction{8, 0}, 5, c, regs},
		0xEA: &setInstruction{basicInstruction{8, 0}, 5, d, regs},
		0xEB: &setInstruction{basicInstruction{8, 0}, 5, e, regs},
		0xEC: &setInstruction{basicInstruction{8, 0}, 5, h, regs},
		0xED: &setInstruction{basicInstruction{8, 0}, 5, l, regs},
		0xEE: &setFromMemoryInstruction{basicInstruction{16, 0}, 5, h, l, regs, mmu},
		0xEF: &setInstruction{basicInstruction{8, 0}, 5, a, regs},
		0xF0: &setInstruction{basicInstruction{8, 0}, 6, b, regs},
		0xF1: &setInstruction{basicInstruction{8, 0}, 6, c, regs},
		0xF2: &setInstruction{basicInstruction{8, 0}, 6, d, regs},
		0xF3: &setInstruction{basicInstruction{8, 0}, 6, e, regs},
		0xF4: &setInstruction{basicInstruction{8, 0}, 6, h, regs},
		0xF5: &setInstruction{basicInstruction{8, 0}, 6, l, regs},
		0xF6: &setFromMemoryInstruction{basicInstruction{16, 0}, 6, h, l, regs, mmu},
		0xF7: &setInstruction{basicInstruction{8, 0}, 6, a, regs},
		0xF8: &setInstruction{basicInstruction{8, 0}, 7, b, regs},
		0xF9: &setInstruction{basicInstruction{8, 0}, 7, c, regs},
		0xFA: &setInstruction{basicInstruction{8, 0}, 7, d, regs},
		0xFB: &setInstruction{basicInstruction{8, 0}, 7, e, regs},
		0xFC: &setInstruction{basicInstruction{8, 0}, 7, h, regs},
		0xFD: &setInstruction{basicInstruction{8, 0}, 7, l, regs},
		0xFE: &setFromMemoryInstruction{basicInstruction{16, 0}, 7, h, l, regs, mmu},
		0xFF: &setInstruction{basicInstruction{8, 0}, 7, a, regs},
	}
}

func (i *swapInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	var flags byte
	flags = 0

	newVal := ((val & 0xF0) >> 4) + ((val & 0x0F) << 4)
	if newVal == 0 {
		flags += 0x10
	}

	i.regs.WriteRegister(i.source, newVal)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *swapFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	newVal := ((val & 0xF0) >> 4) + ((val & 0x0F) << 4)

	var flags byte
	flags = 0
	if newVal == 0 {
		flags += 0x10
	}

	i.mmu.WriteByte(addr, newVal)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rlcInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	var flags byte
	flags = 0

	sevenBit := (val & 0x80) == 0x80
	val = val << 1
	if sevenBit {
		flags += 0x10
		val += 0x01
	}

	i.regs.WriteRegister(i.source, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rlcFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	sevenBit := (val & 0x80) == 0x80
	val = val << 1

	var flags byte
	flags = 0
	if sevenBit {
		flags += 0x10
		val += 0x01
	}

	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rlInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	carryBit := (i.regs.ReadRegister(f) & 0x01) == 0x01
	sevenBit := (val & 0x80) == 0x80
	var flags byte
	flags = 0

	val = val << 1
	if sevenBit {
		flags += 0x10
	}
	if carryBit {
		val += 0x01
	}
	i.regs.WriteRegister(i.source, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rlFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	carryBit := (i.regs.ReadRegister(f) & 0x01) == 0x01
	sevenBit := (val & 0x80) == 0x80
	var flags byte
	flags = 0

	val = val << 1
	if sevenBit {
		flags += 0x10
	}
	if carryBit {
		val += 0x01
	}
	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rrcInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1

	var flags byte
	flags = 0
	if zeroBit {
		flags += 0x10
		val += 0x80
	}

	i.regs.WriteRegister(i.source, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rrcFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1

	var flags byte
	flags = 0
	if zeroBit {
		flags += 0x10
		val += 0x80
	}

	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rrInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	carryBit := (i.regs.ReadRegister(f) & 0x01) == 0x01
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1

	var flags byte
	flags = 0
	if zeroBit {
		flags += 0x10
	}
	if carryBit {
		val += 0x80
	}

	i.regs.WriteRegister(i.source, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *rrFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	carryBit := (i.regs.ReadRegister(f) & 0x01) == 0x01
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1

	var flags byte
	flags = 0
	if zeroBit {
		flags += 0x10
	}
	if carryBit {
		val += 0x80
	}

	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *slaInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	lsb := (val & 0x80) == 0x80
	val = val << 1

	var flags byte
	flags = 0
	if lsb {
		flags += 0x10
	}
	if val == 0 {
		flags += 0x80
	}
	i.regs.WriteRegister(i.source, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *slaFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	lsb := (val & 0x80) == 0x80
	val = val << 1

	var flags byte
	flags = 0
	if lsb {
		flags += 0x10
	}
	if val == 0 {
		flags += 0x80
	}
	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *sraInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	msb := val | 0x80
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1
	val = val + msb

	var flags byte
	flags = 0
	if val == 0 {
		flags += 0x80
	}
	if zeroBit {
		flags += 0x10
	}
	i.regs.WriteRegister(i.source, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *sraFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	msb := val | 0x80
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1
	val = val + msb

	var flags byte
	flags = 0
	if val == 0 {
		flags += 0x80
	}
	if zeroBit {
		flags += 0x10
	}
	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *srlInstruction) Execute(params Parameters) Addresser {
	val := i.regs.ReadRegister(i.source)
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1

	var flags byte
	flags = 0
	if zeroBit {
		flags += 0x10
	}
	if val == 0 {
		flags += 0x80
	}

	i.regs.WriteRegister(i.source, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *srlFromMemoryInstruction) Execute(params Parameters) Addresser {
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}

	val := i.mmu.ReadAt(addr)
	zeroBit := (val & 0x10) == 0x10
	val = val >> 1

	var flags byte
	flags = 0
	if zeroBit {
		flags += 0x10
	}
	if val == 0 {
		flags += 0x80
	}

	i.mmu.WriteByte(addr, val)
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *bitInstruction) Execute(params Parameters) Addresser {
	bit := byte(1 >> i.bitNumber)
	val := i.regs.ReadRegister(i.source)

	var flags byte
	flags = 0x20
	if (val & bit) == 0 {
		flags += 0x80
	}
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *bitFromMemoryInstruction) Execute(params Parameters) Addresser {
	bit := byte(1 >> i.bitNumber)
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)

	var flags byte
	flags = 0x20
	if (val & bit) == 0 {
		flags += 0x80
	}
	i.regs.WriteRegister(f, flags)
	return &address{}
}

func (i *setInstruction) Execute(params Parameters) Addresser {
	bit := byte(1 >> i.bitNumber)
	i.regs.WriteRegister(i.source, (i.regs.ReadRegister(i.source) | bit))
	return &address{}
}

func (i *setFromMemoryInstruction) Execute(params Parameters) Addresser {
	bit := byte(1 >> i.bitNumber)
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	i.mmu.WriteByte(addr, val|bit)
	return &address{}
}

func (i *resetInstruction) Execute(params Parameters) Addresser {
	bit := byte(1>>i.bitNumber) ^ 0xFF
	i.regs.WriteRegister(i.source, (i.regs.ReadRegister(i.source) & bit))
	return &address{}
	return &address{}
}

func (i *resetFromMemoryInstruction) Execute(params Parameters) Addresser {
	bit := byte(1>>i.bitNumber) ^ 0xFF
	addr, err := i.regs.ReadRegisterPair(i.source1, i.source2)
	if err != nil {
		fmt.Println(err)
	}
	val := i.mmu.ReadAt(addr)
	i.mmu.WriteByte(addr, val&bit)
	return &address{}
}
