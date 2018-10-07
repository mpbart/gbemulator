package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

const (
	CLOCK_TICK_NANOS = (time.Second / (2 ^ 35)) * time.Nanosecond
)

func initialize() {
	runtime.LockOSThread()
	//runtime.GOMAXPROCS(runtime.NumCPU() - 1)
}

func main() {
	exitChannel := make(chan bool)

	mmu := InitializeMMU()
	cpu := InitializeCPU(exitChannel, mmu)

	CreateDisplay(mmu, cpu) // GLFW wil not work if the window pointer is passed around so this function never returns
}

func InitializeCPU(exitChannel chan bool, mmu MMU) CPU {
	cpu := CreateCPU(exitChannel, mmu)
	cpu.Reset()
	return cpu
}

func InitializeMMU() MMU {
	mmu := CreateMMU()
	mmu.Reset()

	if err := loadROM(mmu); err != nil {
		fmt.Println(err)
		return nil
	}
	return mmu
}

func loadROM(m MMU) error {
	// Hardcode ROM name for now, maybe add in some command line arguments to parse this later
	filename := "tetris.gb"
	f, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("ERROR opening ROM: %s", err)
	}

	for i := 0; i < len(f); i++ {
		m.WriteByte(uint16(i), f[i])
	}
	return nil
}
