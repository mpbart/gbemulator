package main

import "fmt"

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

func CreateExtendedInstructions(regs Registers, mmu MMU) map[byte]ExtendedInstruction {
	return map[byte]ExtendedInstruction{
		0x30: &swapInstruction{basicInstruction{8, 0}, b, regs},
		0x31: &swapInstruction{basicInstruction{8, 0}, c, regs},
		0x32: &swapInstruction{basicInstruction{8, 0}, d, regs},
		0x33: &swapInstruction{basicInstruction{8, 0}, e, regs},
		0x34: &swapInstruction{basicInstruction{8, 0}, h, regs},
		0x35: &swapInstruction{basicInstruction{8, 0}, l, regs},
		0x36: &swapFromMemoryInstruction{basicInstruction{16, 0}, h, l, regs, mmu},
		0x37: &swapInstruction{basicInstruction{8, 0}, a, regs},
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
