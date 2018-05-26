package cpu

type 8BitRegister byte
type 16BitRegister uint16

type GbMachineState struct {
	ROM  []uint8
	RAM  []uint8
	VRAM []uint8

	PC   16BitRegister // Should be initialized to 0x100 to start execution
	SP	 16BitRegister // Should be initialized to 0xFFFE on startup (grows downward in RAM)

	RegA 8BitRegister
	RegB 8BitRegister
	RegC 8BitRegister
	RegD 8BitRegister
	RegE 8BitRegister
	RegF 8BitRegister
	RegH 8BitRegister
	RegL 8BitRegister
}
