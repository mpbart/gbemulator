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
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
}

func main() {
	exitChannel := make(chan bool)
	gpu := CreateDisplay()
	go CreateComponents(exitChannel, gpu)
	gpu.Start(exitChannel) // This needs to happen on the main OS thread for the UI library to function correctly
}

func CreateComponents(exitChannel chan bool, gpu Display) {
	m := CreateMMU()
	cpu := CreateCPU(m, gpu)

	m.Reset()
	cpu.Reset()
	if err := loadROM(m); err != nil {
		fmt.Println(err)
		return
	}
	cpu.Run(exitChannel)
}

func loadROM(m MMU) error {
	// Hardcode ROM name for now, maybe add in some command line arguments to parse this later
	filename := "bgbtest.rom"
	f, err := ioutil.ReadFile(filename)

	if err != nil {
		return fmt.Errorf("ERROR opening ROM: %s", err)
	}

	for i := 0; i < len(f); i++ {
		m.WriteByte(uint16(i), f[i])
	}
	return nil
}
