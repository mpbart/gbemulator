package main

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	SCREEN_WIDTH  int = 160
	SCREEN_HEIGHT int = 144
)

type display struct {
	window      *glfw.Window
	mmu         MMU
	tickChannel chan bool
}

type RGBPixel struct {
	Red   uint
	Green uint
	Blue  uint
}

type Display interface {
	Render()
	Tick()
	Start()
}

// Notes:
// * 4 Modes are
//		- 1. OAM - 20 clocks
//			* search for sprites that are visible on this line
//			* oam.x != 0 && LY + 16 >= oam.y && LY + 16 < oam.y
//		- 2. Pixel transfer - 43+ clocks (can be more depending on window and pixels drawn)
//			* Shifts out one pixel per (4 mHZ) clock from the PPU
//			* Needs to also store the source of the pixel to determine priority
//			* Lower numbered sprites > higher sprites > background
//			* Fetches the next 8 pixels at half speed
//			* Fetches take 4 steps
//				- 1. Read tile number
//				- 2. Read byte 1
//				- 3. Read byte 2
//				- 4. Idle until the last 8 bits in the PPU are empty
//			* Windows cause the PPU to be totally reset and to start fetching from that window location
//			* When a sprite is encountered it's pixels are overlaid onto the first 8 pixels in the PPU
//		- 3. H-Blank - 51- clocks (Extra clocks in pixel transfer are taken out of H-Blank)
//		... Repeats 144 times
//		- 4. V-Blank - (20 + 43 + 51) clocks
//		... Repeats 10 times
//		- 5. Goes back to mode 1

// The 4 methods below are intended to be used as constants
func WHITE() RGBPixel {
	return RGBPixel{0, 0, 0}
}

func LIGHT_GRAY() RGBPixel {
	return RGBPixel{211, 211, 211}
}

func DARK_GRAY() RGBPixel {
	return RGBPixel{47, 70, 79}
}

func BLACK() RGBPixel {
	return RGBPixel{0, 0, 0}
}

func CreateDisplay(mmu MMU, cpu CPU) {
	err := glfw.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		fmt.Println(err)
	}

	glfw.WindowHint(glfw.OPENGL_FORWARD_COMPAT, GL_TRUE)
	window, err := glfw.CreateWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "GB Emulator", nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	window.MakeContextCurrent()

	gl.Viewport(0, 0, int32(SCREEN_WIDTH), int32(SCREEN_HEIGHT))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(SCREEN_WIDTH), float64(SCREEN_HEIGHT), 0, -1, 1)
	gl.ClearColor(0.255, 0.255, 0.255, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	window.SetPos(0, 0)

	tickChannel := make(chan bool)
	d := &display{window, mmu, tickChannel}
	d.Simulate(cpu, mmu)
}

func (d *display) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Disable(gl.DEPTH_TEST)
	gl.PointSize(1.0)
	gl.Begin(gl.POINTS)
	for y := 0; y < SCREEN_HEIGHT; y++ {
		for x := 0; x < SCREEN_WIDTH; x++ {
			if x%4 == 0 {
				gl.Color3ub(30, 20, 10)
				gl.Vertex2i(int32(x), int32(y))
			}
		}
	}
	gl.End()
	d.window.SwapBuffers()
}

func (d *display) Simulate(cpu CPU, mmu MMU) {
	ticks := 0
	for {
		cpu.Tick()
		mmu.Tick()
		if ticks >= 70224 {
			d.Tick()
			ticks = 0
		}
		ticks += 1
	}
}

func (d *display) Tick() {
	if d.window.ShouldClose() {
		return
	}
	d.Render()
	d.window.SwapBuffers()
	glfw.PollEvents()
}

func (d *display) mode() uint8 {
	return d.mmu.LCDStatusMode()
}
