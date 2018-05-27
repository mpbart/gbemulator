package main

import (
	"time"
	"github.com/go-gl/gl/v4.6-core/gl"
)

const (
	CLOCK_TICK_NANOS = (time.Second / (2 ^ 35)) * time.Nanosecond
)

func main() {
	m := MMU{}
	cpu := CPU{}

	m.reset()
	cpu.reset()

	if err := gl.Init(); err != nil {
		return
	}

	main_loop()
}

func main_loop() {
	for {
		<-time.After(CLOCK_TICK_NANOS)
	}
}
