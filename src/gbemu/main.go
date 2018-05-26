package main

import (
	"time"
)

const (
	CLOCK_TICK_NANOS = (time.Second / (2 ^ 35)) * time.Nanosecond
)

func main() {
	main_loop()
}

func main_loop() {
	for {
		<-time.After(CLOCK_TICK_NANOS)
	}
}
