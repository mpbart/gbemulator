package main

type PPU interface {
	Tick([]SpriteAttribute, int)
	LineFinished() bool
	Reset()
	LcdBuffer(int, int) RGBPixel
}

type ppu struct {
	fifo                      []RGBPixel
	fetcher                   Fetcher
	lcdBuffer                 [][]RGBPixel
	currentPixel              uint16
	currentSpritePixelFetched bool
	fetchingSprite            bool
}

func createPPU(mmu MMU) PPU {
	buffer := make([][]RGBPixel, SCREEN_HEIGHT)
	for i := range buffer {
		buffer[i] = make([]RGBPixel, SCREEN_WIDTH)
	}

	return &ppu{
		fifo:                      make([]RGBPixel, 0),
		fetcher:                   createFetcher(mmu),
		currentPixel:              0,
		currentSpritePixelFetched: false,
		lcdBuffer:                 buffer,
		fetchingSprite:            false,
	}
}

func (p *ppu) Tick(sprites []SpriteAttribute, currentLine int) {
	// Shifts in 8 pixels at a time from the fetcher, if they are available
	if pixels := p.fetcher.Fetch(currentLine); pixels != nil {
		if p.fetchingSprite {
			p.currentSpritePixelFetched = true
			p.overlayPixels(pixels)
			p.fetchingSprite = false
			p.fetcher.Reset(uint16(p.currentPixel), BG_FETCH, nil)
		} else {
			p.shiftInPixels(pixels, currentLine)
		}
	}

	if p.canShift() {
		p.shiftOutPixel(currentLine, sprites)
	}
}

func (p *ppu) Reset() {
	p.currentPixel = 0
}

func (p *ppu) canShiftOut() bool {
	return len(p.fifo) > 8 && p.currentPixel < uint16(SCREEN_WIDTH) && !p.fetchingSprite
}

func (p *ppu) LineFinished() bool {
	return p.currentPixel == uint16(SCREEN_WIDTH)
}

// Will fail when sprite finishes fetching and ppu attempts to shift out next pixel because it will see the same
// sprite and think it needs to re-fetch...
func (p *ppu) shiftOutPixel(currentLine int, sprites []SpriteAttribute) {
	if !p.isUnfetchedSpritePixel(sprites) {
		p.lcdBuffer[currentLine][p.currentPixel] = p.fifo[0]
		p.fifo = p.fifo[1:]
		p.currentPixel += 1
		p.currentSpritePixelFetched = false
	}
}

func (p *ppu) isUnfetchedSpritePixel(sprites []SpriteAttribute) bool {
	for _, sprite := range sprites {
		if sprite != nil && sprite.GetXPosition() > 0 && sprite.GetXPosition()-8 == int(p.currentPixel) && !p.currentSpritePixelFetched {
			p.fetcher.Reset(uint16(p.currentPixel), SPRITE_FETCH, sprite)
			p.fetchingSprite = true
			return true
			// 1. Stop current fetching
			// 2. Start fetching sprite pixels
			// 3. Overlay first 8 pixels from background with sprite pixels
			// 4. Resume fetching background
		}
	}
	return false
}

func (p *ppu) shiftInPixels(pixels []RGBPixel, currentLine int) {
	for _, pixel := range pixels {
		p.fifo = append(p.fifo, pixel)
	}
}

func (p *ppu) LcdBuffer(y, x int) RGBPixel {
	return p.lcdBuffer[y][x]
}

func (p *ppu) overlayPixels(pixels []RGBPixel) {
}
