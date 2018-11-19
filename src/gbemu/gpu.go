package main

import (
	"fmt"
	"sort"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	SCREEN_WIDTH        int   = 160
	SCREEN_HEIGHT       int   = 144
	HBLANK_MODE         uint8 = 0
	VBLANK_MODE         uint8 = 1
	OAM_SEARCH_MODE     uint8 = 2
	PIXEL_TRANSFER_MODE uint8 = 3
)

type display struct {
	window         *glfw.Window
	mmu            MMU
	ppu            PPU
	addresser      MemoryAddresser
	currentTicks   int
	lY             int
	visibleSprites []SpriteAttribute
}

type RGBPixel struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Display interface {
	Render()
	Tick()
	Start()
	CurrentLine()
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
	return RGBPixel{255, 255, 255}
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

func CreateDisplay(mmu MMU, cpu CPU, timer Timer) {
	err := glfw.Init()
	if err != nil {
		fmt.Println(err)
	}
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		fmt.Println(err)
	}

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

	d := &display{
		window:         window,
		mmu:            mmu,
		addresser:      CreateMemoryAddresser(mmu.BGAndWindowAddressMode()),
		ppu:            createPPU(mmu),
		currentTicks:   0,
		lY:             0,
		visibleSprites: make([]SpriteAttribute, 10),
	}
	d.Simulate(cpu, mmu, timer)
}

func (d *display) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Disable(gl.DEPTH_TEST)
	gl.PointSize(1.0)
	gl.Begin(gl.POINTS)
	for y := 0; y < SCREEN_HEIGHT; y++ {
		for x := 0; x < SCREEN_WIDTH; x++ {
			pixel := d.ppu.LcdBuffer(y, x)
			gl.Color3ub(pixel.Red, pixel.Green, pixel.Blue)
			gl.Vertex2i(int32(x), int32(y))
		}
	}
	gl.End()
	d.window.SwapBuffers()
}

func (d *display) Simulate(cpu CPU, mmu MMU, timer Timer) {
	for {
		cpu.Tick()
		mmu.Tick()
		timer.Tick()
		d.Tick()
	}
}

func (d *display) Tick() {
	if !d.mmu.LCDEnabled() {
		return
	}

	switch d.mode() {
	case OAM_SEARCH_MODE:
		if d.currentTicks == 80 {
			d.readOam()
			d.mmu.SetLCDStatusMode(PIXEL_TRANSFER_MODE)
			d.currentTicks = 0
		} else {
			d.currentTicks += 1
		}
	case PIXEL_TRANSFER_MODE:
		d.ppu.Tick(d.visibleSprites, d.lY)
		if d.ppu.LineFinished() {
			d.mmu.SetLCDStatusMode(HBLANK_MODE)
			d.ppu.Reset()
		}
		d.currentTicks += 1
	case HBLANK_MODE:
		if d.currentTicks == 376 {
			if d.lY == SCREEN_HEIGHT-1 {
				d.mmu.SetLCDStatusMode(VBLANK_MODE)
				d.lY += 1
				d.mmu.WriteByte(LCDC_Y_COORDINATE, uint8(d.lY))
				d.currentTicks = 0
				d.mmu.FireInterrupt(VBLANK_INTERRUPT)
			} else {
				d.lY += 1
				d.mmu.WriteByte(LCDC_Y_COORDINATE, uint8(d.lY))
				d.mmu.SetLCDStatusMode(OAM_SEARCH_MODE)
				d.currentTicks = 0
			}
		} else {
			d.currentTicks += 1
		}
	case VBLANK_MODE:
		if d.currentTicks == 456 {
			if d.lY == SCREEN_WIDTH+9 {
				d.mmu.SetLCDStatusMode(OAM_SEARCH_MODE)
				d.currentTicks = 0
				d.lY = 0
				d.updateDisplay()
			} else {
				d.lY += 1
				d.mmu.WriteByte(LCDC_Y_COORDINATE, uint8(d.lY))
				d.currentTicks = 0
			}
		} else {
			d.currentTicks += 1
		}
	}
}

func (d *display) updateDisplay() {
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

func (d *display) readOam() {
	d.visibleSprites = nil
	for i := 0; i < 40; i++ {
		byte0 := d.mmu.ReadAt(uint16(0xFE00 + i*4))
		byte1 := d.mmu.ReadAt(uint16(0xFE00 + i*4 + 1))
		byte2 := d.mmu.ReadAt(uint16(0xFE00 + i*4 + 2))
		byte3 := d.mmu.ReadAt(uint16(0xFE00 + i*4 + 3))
		attr := fromBytes([]uint8{byte0, byte1, byte2, byte3})

		if d.lY >= attr.GetYPosition()-16 && d.lY < attr.GetYPosition()-16+d.spriteHeight() {
			d.visibleSprites = append(d.visibleSprites, attr)
		}

		// Only Store the first 10 sprites to be rendered
		if len(d.visibleSprites) == 10 {
			break
		}
	}

	sort.Stable(SortableSpriteAttribute(d.visibleSprites))
}

func (d *display) spriteHeight() int {
	sizeBit := d.mmu.SpriteSize()
	if sizeBit == 0 {
		return 8
	} else {
		return 16
	}
}

func (d *display) CurrentLine() int {
	return d.lY
}
