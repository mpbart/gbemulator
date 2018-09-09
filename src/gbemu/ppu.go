package main

type PPU interface {
	Tick([]SpriteAttribute, int)
	LineFinished() bool
}

type ppu struct {
	fifo         []RGBPixel
	fetcher      Fetcher
	lcdBuffer    [][]RGBPixel
	currentPixel int
}

func createPPU(mmu MMU) PPU {
	buffer := make([][]RGBPixel, SCREEN_HEIGHT)
	for i := range buffer {
		buffer[i] = make([]RGBPixel, SCREEN_WIDTH)
	}

	return &ppu{
		fifo:         make([]RGBPixel, 16),
		fetcher:      createFetcher(mmu),
		currentPixel: 0,
		lcdBuffer:    buffer,
	}
}

func (p *ppu) Tick(_ []SpriteAttribute, currentLine int) {
	if pixels := p.fetcher.Fetch(); pixels != nil {
		p.shiftInPixels(pixels, currentLine)
	}

	if p.canShift() {
		p.shiftOutPixel(currentLine)
	}
}

func (p *ppu) canShift() bool {
	return len(p.fifo) > 8
}

func (p *ppu) LineFinished() bool {
	return false
}

func (p *ppu) shiftOutPixel(currentLine int) {
}

func (p *ppu) shiftInPixels(pixels []RGBPixel, currentLine int) {
	// If FIFO is empty then put pixels at beginning of fifo, otherwise add them to the end
	start := 0
	if len(p.fifo) != 0 {
		start = 8
	}

	for i := 0 + start; i < len(pixels)+start; i++ {
		p.fifo[i] = pixels[i]
	}
}
