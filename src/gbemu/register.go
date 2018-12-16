package main

import "fmt"

type Register int

const (
	a Register = iota
	b
	c
	d
	e
	f
	h
	l
)

// combinations of 8 bit registers that can be combined to form 16 bit registers
var validRegisterPairs = [][]Register{
	[]Register{a, f},
	[]Register{b, c},
	[]Register{d, e},
	[]Register{h, l},
}

type Registers interface {
	ReadRegister(Register) byte
	WriteRegister(Register, byte)
	ReadRegisterPair(Register, Register) (uint16, error)
	WriteRegisterPair(Register, Register, uint16) error
	DumpContents(int)

	ReadPC() uint16
	WritePC(uint16)
	PopSP() uint16
	PushSP(uint16)
	WriteSP(uint16)
	ReadSP() uint16
}

type registers struct {
	PC uint16 // Should be initialized to 0x100 to start execution
	SP uint16 // Should be initialized to 0xFFFE on startup (grows downward in RAM)

	mmu MMU

	regs map[Register]byte
}

func CreateRegisters(mmu MMU) Registers {
	r := registers{
		PC:   0,
		SP:   0,
		mmu:  mmu,
		regs: make(map[Register]byte),
	}

	r.regs[a] = 0
	r.regs[b] = 0
	r.regs[c] = 0
	r.regs[d] = 0
	r.regs[e] = 0
	r.regs[f] = 0
	r.regs[h] = 0
	r.regs[l] = 0

	return &r
}

func (r *registers) PopSP() uint16 {
	r.SP += 0x02
	lsb := r.mmu.ReadAt(r.SP)
	msb := r.mmu.ReadAt(r.SP + 1)
	return (uint16(msb) << 8) + uint16(lsb)
}

func (r *registers) PushSP(value uint16) {
	r.mmu.WriteByte(r.SP, byte(value&0xFF))
	r.mmu.WriteByte(r.SP+1, byte(value>>8))
	r.SP -= 0x02
}

func (r *registers) ReadSP() uint16 {
	return r.SP
}

func (r *registers) WriteSP(value uint16) {
	r.SP = value
}

func (r *registers) ReadPC() uint16 {
	return r.PC
}

func (r *registers) WritePC(value uint16) {
	r.PC = value
}

func (r *registers) ReadRegister(reg Register) byte {
	return r.regs[reg]
}

func (r *registers) WriteRegister(reg Register, value byte) {
	r.regs[reg] = value
}

func (r *registers) ReadRegisterPair(reg1, reg2 Register) (retval uint16, err error) {
	if !r.validRegisterPair(reg1, reg2) {
		err = fmt.Errorf("Invalid register Pair: %v, %v", reg1, reg2)
		return
	}
	retval = (uint16(r.regs[reg1]) << 8) + uint16(r.regs[reg2])
	return
}

func (r *registers) WriteRegisterPair(reg1, reg2 Register, value uint16) error {
	if !r.validRegisterPair(reg1, reg2) {
		err := fmt.Errorf("Invalid register Pair: %v, %v", reg1, reg2)
		return err
	}
	r.regs[reg1] = byte(value >> 8)
	r.regs[reg2] = byte(value & 0xFF)
	return nil
}

func (r *registers) validRegisterPair(reg1, reg2 Register) bool {
	regPair := []Register{reg1, reg2}
	for i := 0; i < len(validRegisterPairs); i++ {
		if validRegisterPairs[i][0] == regPair[0] && validRegisterPairs[i][1] == regPair[1] {
			return true
		}
	}
	return false
}

func (r *registers) GetValue() byte {
	return 0
}

func (r *registers) GetTwoByteValue() uint16 {
	return 0
}

func (r *registers) DumpContents(step int) {
	//fmt.Printf("Register values:\na: %x\nb: %x\nc: %x\nd: %x\ne: %x\nh: %x\nl: %x\nf: %x\nsp: %x\npc: %x\n", r.regs[a], r.regs[b], r.regs[c], r.regs[d], r.regs[e], r.regs[h], r.regs[l], r.regs[f], r.SP, r.PC)
	af, _ := r.ReadRegisterPair(a, f)
	bc, _ := r.ReadRegisterPair(b, c)
	de, _ := r.ReadRegisterPair(d, e)
	hl, _ := r.ReadRegisterPair(h, l)
	fmt.Printf("Step:%08d, ", step)
	fmt.Printf("[PC:%04x ", r.PC)
	fmt.Printf("SP:%04x ", r.SP)
	fmt.Printf("AF:%04x ", af)
	fmt.Printf("BC:%04x ", bc)
	fmt.Printf("DE:%04x ", de)
	fmt.Printf("HL:%04x ", hl)
	fmt.Printf("LY:%02x ", r.mmu.ReadAt(LCDC_Y_COORDINATE))
	fmt.Printf("LS:%02x \n", r.mmu.ReadAt(LCDC_STATUS))
}
