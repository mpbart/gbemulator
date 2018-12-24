package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Currently Investigating:
//	* tile addressing scheme changes at 0x0239 or somewhere before it. Does not look like
//    we are changing with it because we are not clearing the screen correctly.

const TICKS_PER_REFRESH int = 70224

type CPU interface {
	Reset()
	Tick()
	SetInterruptMasterEnable(bool)
}

type cpu struct {
	mmu                   MMU
	registers             Registers
	instructions          map[byte]Instruction
	stopped               bool
	InstructionTicks      int
	exitChannel           chan bool
	currentInstruction    Instruction
	currentOpcode         byte
	currentParamBytes     int
	currentParams         Parameters
	getInput              bool
	interruptMasterEnable bool
	ticks                 int
	breakAddresses        []uint16
}

func CreateCPU(exitChannel chan bool, mmu MMU) CPU {
	registers := CreateRegisters(mmu)

	cpu := &cpu{
		mmu:                   mmu,
		registers:             registers,
		instructions:          nil,
		stopped:               false,
		InstructionTicks:      4, // Current assumption is that the first instruction is always 4 cycles. May need to refactor and hold instructions as state to fix this
		exitChannel:           exitChannel,
		currentInstruction:    nil,
		currentOpcode:         0,
		currentParamBytes:     0,
		currentParams:         Parameters{},
		getInput:              false,
		interruptMasterEnable: false,
		ticks:                 0,
		breakAddresses:        []uint16{0x02BC},
	}

	cpu.instructions = CreateInstructions(registers, mmu, cpu)
	return cpu
}

func (cpu *cpu) Reset() {
	cpu.registers.WritePC(0x100)
	cpu.registers.PushSP(0xFFFE)
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
	cpu.mmu.WriteByte(0xFF41, 0x02) // Set status to OAM
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

	cpu.stopped = false
	cpu.decodeNextInstruction()
}

func (c *cpu) IncrementPC(offset int) {
	c.registers.WritePC(c.registers.ReadPC() + uint16(offset))
}

func (c *cpu) Tick() {
	if c.mmu.HasPendingInterrupt() && c.interruptMasterEnable {
		c.executeInstruction() // Always execute the pending instruction before running interrupt routine

		interruptVector := c.mmu.GetNextPendingInterrupt()
		c.mmu.ClearHighestInterrupt()
		c.registers.PushSP(c.registers.ReadPC() + uint16(c.currentInstruction.GetNumParameterBytes()))
		fmt.Printf("Pushing PC: %x\n", c.registers.ReadPC())
		c.registers.WritePC(interruptVector)
		c.InstructionTicks = 12
		c.interruptMasterEnable = false

		c.decodeNextInstruction() // Get and decode instruction at the interrupt
	}

	if c.InstructionTicks != 0 {
		c.InstructionTicks -= 1
		return
	} else {
		// Don't execute any more instructions until a key press event happens
		if c.stopped {
			return
		}
		c.executeInstruction()
		c.decodeNextInstruction()
	}
}

func (c *cpu) getNextInstruction() uint8 {
	return c.mmu.ReadAt(c.registers.ReadPC())
}

func (c *cpu) decodeNextInstruction() {
	c.currentOpcode = c.getNextInstruction()
	instruction, found := c.instructions[c.currentOpcode]

	if !found {
		fmt.Printf("ERROR: Opcode %x not found\n", c.currentOpcode)
		c.exitChannel <- true
		return
	}

	c.currentInstruction = instruction
	c.currentParamBytes = c.currentInstruction.GetNumParameterBytes()
	c.currentParams = make(Parameters, c.currentParamBytes)
	if c.currentParamBytes > 0 {
		for i := 0; i < c.currentParamBytes; i++ {
			c.currentParams[i] = c.mmu.ReadAt(c.registers.ReadPC() + uint16(i+1))
		}
	}

	c.InstructionTicks = instruction.GetCycles(c.currentParams) // TODO: Need a better way of doing this. Pretty awkward to send in params right now...
}

func (c *cpu) executeInstruction() {
	result := c.currentInstruction.Execute(c.currentParams)

	//fmt.Printf("executed %x at %x\n", c.currentOpcode, c.registers.ReadPC())
	if c.atBreakpoint() || c.getInput {
		c.registers.DumpContents(c.ticks)
	debug:
		for {
			fmt.Print("<Enter Debug Command> ")
			reader := bufio.NewReader(os.Stdin)
			str, _ := reader.ReadString('\n')
			command := strings.Trim(str, "\n")
			switch command {
			case "M", "m": //Check memory location
				fmt.Print("Enter memory location to print:")
				rawAddress, _ := reader.ReadString('\n')
				address := strings.Trim(rawAddress, "\n")
				addr, err := strconv.ParseUint(address, 16, 16)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Printf("%x\n", c.mmu.ReadAt(uint16(addr)))
			case "C", "c": // continue
				c.getInput = true
				break debug
			case "E", "e": // exit debugger and continue
				c.getInput = false
				break debug
			/*
				case "R", "r": // run to address and then break
					fmt.Print("Enter breakpoint to remove")
					rawPC, _ := reader.ReadString('\n')
					strippedPC := strings.Trim(rawPC, "\n")
					pc, err := strconv.ParseUint(strippedPC, 16, 16)
					if err != nil {
						fmt.Println(err)
						continue
					}
					c.breakAddress = uint16(pc)
			*/
			case "B", "b": // Add breakpoint
				c.getInput = false
				fmt.Print("Enter PC to break at:")
				rawPC, _ := reader.ReadString('\n')
				strippedPC := strings.Trim(rawPC, "\n")
				pc, err := strconv.ParseUint(strippedPC, 16, 16)
				if err != nil {
					fmt.Println(err)
					continue
				}
				c.breakAddresses = append(c.breakAddresses, uint16(pc))
			case "H", "h": // print help
				fmt.Println("M,m - view value at memory location 0x<input>")
				fmt.Println("C,c - continue and break after next cpu cycle")
				fmt.Println("E,e - exit debugger and continue running")
				fmt.Println("B,b - add new breakpoint")
			default:
				continue
			}
		}
	}

	if result.ShouldJump() {
		c.registers.WritePC(result.NewAddress())
	} else {
		c.IncrementPC(c.currentParamBytes + 1)
	}

	if result.IsStopped() {
		c.stopped = true
	}
	c.ticks += 1
}

func (c *cpu) SetInterruptMasterEnable(value bool) {
	c.interruptMasterEnable = value
}

func (c *cpu) atBreakpoint() bool {
	currentPc := c.registers.ReadPC()
	for _, addr := range c.breakAddresses {
		if addr == currentPc {
			return true
		}
	}
	return false
}
