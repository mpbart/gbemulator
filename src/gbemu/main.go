package main

import (
	"time"
	"runtime"
)

const (
	CLOCK_TICK_NANOS = (time.Second / (2 ^ 35)) * time.Nanosecond
)

func initialize() {
	runtime.LockOSThread()
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
}

func main() {
	go CreateComponents()
	CreateDisplay() // This needs to happen on the main OS thread for the UI library to function correctly
}

func CreateComponents() {
	m := CreateMMU()
	cpu := CreateCPU(m)

	m.Reset()
	cpu.Reset()
}
