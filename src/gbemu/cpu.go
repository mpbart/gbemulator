package main

import "fmt"

type CPU interface {
	Reset()
	Run(chan bool)
}

type cpu struct {
	mmu          MMU
	registers    Registers
	instructions map[byte]Instruction
}

func CreateCPU(mmu MMU) CPU {
	registers := CreateRegisters(mmu)

	return &cpu{
		mmu:          mmu,
		registers:    registers,
		instructions: CreateInstructions(registers, mmu),
	}
}

func (cpu *cpu) Reset() {
	cpu.registers.WritePC(0x100)
	cpu.registers.WriteSP(0xFFFE)
	cpu.registers.WriteRegisterPair(a, f, 0x01B0)
	cpu.registers.WriteRegisterPair(b, c, 0x0013)
	cpu.registers.WriteRegisterPair(d, e, 0x00D8)
	cpu.registers.WriteRegisterPair(h, l, 0x014D)

	cpu.mmu.WriteByte(0xFF05, 0x00)
	cpu.mmu.WriteByte(0xFF06, 0x00)
	cpu.mmu.WriteByte(0xFF07, 0x00)
	cpu.mmu.WriteByte(0xFF10, 0x80)
	cpu.mmu.WriteByte(0xFF11, 0xBF)
	cpu.mmu.WriteByte(0xFF12, 0xF3)
	cpu.mmu.WriteByte(0xFF14, 0xBF)
	cpu.mmu.WriteByte(0xFF16, 0x3F)
	cpu.mmu.WriteByte(0xFF17, 0x00)
	cpu.mmu.WriteByte(0xFF19, 0xBF)
	cpu.mmu.WriteByte(0xFF1A, 0x7F)
	cpu.mmu.WriteByte(0xFF1B, 0xFF)
	cpu.mmu.WriteByte(0xFF1C, 0x9F)
	cpu.mmu.WriteByte(0xFF1E, 0xBF)
	cpu.mmu.WriteByte(0xFF20, 0xFF)
	cpu.mmu.WriteByte(0xFF21, 0x00)
	cpu.mmu.WriteByte(0xFF22, 0x00)
	cpu.mmu.WriteByte(0xFF23, 0xBF)
	cpu.mmu.WriteByte(0xFF24, 0x77)
	cpu.mmu.WriteByte(0xFF25, 0xF3)
	cpu.mmu.WriteByte(0xFF26, 0xF1)
	cpu.mmu.WriteByte(0xFF40, 0x91)
	cpu.mmu.WriteByte(0xFF42, 0x00)
	cpu.mmu.WriteByte(0xFF43, 0x00)
	cpu.mmu.WriteByte(0xFF45, 0x00)
	cpu.mmu.WriteByte(0xFF47, 0xFC)
	cpu.mmu.WriteByte(0xFF48, 0xFF)
	cpu.mmu.WriteByte(0xFF49, 0xFF)
	cpu.mmu.WriteByte(0xFF4A, 0x00)
	cpu.mmu.WriteByte(0xFF4B, 0x00)
	cpu.mmu.WriteByte(0xFF50, 0x00)
	cpu.mmu.WriteByte(0xFFFF, 0x00)
}

func (c *cpu) IncrementSP() {
	c.registers.WriteSP(c.registers.ReadSP() + 0x02)
}

func (c *cpu) DecrementSP() {
	c.registers.WriteSP(c.registers.ReadSP() - 0x02)
}

func (c *cpu) IncrementPC(offset int) {
	c.registers.WritePC(c.registers.ReadPC() + uint16(offset))
}

func (c *cpu) Run(exitChannel chan bool) {
	for {
		opcode := c.getNextInstruction()
		instruction, found := c.instructions[opcode]

		if !found {
			fmt.Printf("ERROR: Opcode %x not found\n", opcode)
			exitChannel <- true
			break
		}

		paramBytes := instruction.GetNumParameterBytes()
		params := make(Parameters, paramBytes)
		if paramBytes > 0 {
			for i := 0; i < paramBytes; i++ {
				params[i] = c.mmu.ReadAt(c.registers.ReadPC() + uint16(i+1))
			}
		}
		result := instruction.Execute(params)

		fmt.Printf("executed %x\n", opcode)

		if result.ShouldJump() {
			c.registers.WritePC(result.NewAddress())
		} else {
			c.IncrementPC(paramBytes + 1)
		}
	}
}

func (c *cpu) getNextInstruction() byte {
	return c.mmu.ReadAt(c.registers.ReadPC())
}
