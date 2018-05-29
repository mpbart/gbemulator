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
}

func main() {
	go CreateComponents()
	CreateDisplay()
}

func CreateComponents() {
	m := MMU{}
	cpu := CPU{}

	m.reset()
	cpu.reset()
}
