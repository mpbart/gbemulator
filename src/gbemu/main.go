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
	timer := InitializeTimer(mmu)

	CreateDisplay(mmu, cpu, timer) // GLFW wil not work if the window pointer is passed around so this function never returns
}

func InitializeCPU(exitChannel chan bool, mmu MMU) CPU {
	cpu := CreateCPU(exitChannel, mmu)
	cpu.Reset()
	return cpu
}

func InitializeTimer(mmu MMU) Timer {
	timer := CreateTimer(mmu)
	return timer
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

	// Setup displaying of Nintendo logo
	lookupTable := []byte{
		0x00, 0x03, 0x0c, 0x0f, 0x30, 0x33, 0x3c, 0x3f,
		0xc0, 0xc3, 0xcc, 0xcf, 0xf0, 0xf3, 0xfc, 0xff,
	}

	hdrTileData := []byte{}
	for i := 0x104; i < 0x104+48; i++ {
		value := m.ReadAt(uint16(i))
		v1, v2 := lookupTable[value>>4], lookupTable[value&0x0f]
		hdrTileData = append(hdrTileData, v1, 0, v1, 0, v2, 0, v2, 0)
	}

	hdrTileData = append(hdrTileData,
		0x3c, 0x00, 0x42, 0x00, 0xb9, 0x00, 0xa5, 0x00, 0xb9, 0x00, 0xa5, 0x00, 0x42, 0x00, 0x3c, 0x00,
	)

	bootTileMap := []byte{
		0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c,
		0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x0d, 0x0e, 0x0f, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
	}

	for i := range hdrTileData {
		m.WriteByte(uint16(0x8010+i), hdrTileData[i])
	}
	for i := range bootTileMap {
		m.WriteByte(uint16(0x9900+i), bootTileMap[i])
	}
	return nil
}
