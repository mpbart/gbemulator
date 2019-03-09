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
	currentFetchPixel         uint16
	lcdCurrentPixel           uint16
	currentSpritePixelFetched bool
	fetchingSprite            bool
	mmu                       MMU
	lastFetchedSprite         SpriteAttribute
}

func createPPU(mmu MMU) PPU {
	buffer := make([][]RGBPixel, SCREEN_HEIGHT)
	for i := range buffer {
		buffer[i] = make([]RGBPixel, SCREEN_WIDTH)
	}

	return &ppu{
		fifo:                      make([]RGBPixel, 0),
		fetcher:                   createFetcher(mmu),
		currentFetchPixel:         0,
		lcdCurrentPixel:           0,
		currentSpritePixelFetched: false,
		lcdBuffer:                 buffer,
		fetchingSprite:            false,
		mmu:                       mmu,
	}
}

func (p *ppu) Tick(sprites []SpriteAttribute, currentLine int) {
	// Shifts in 8 pixels at a time from the fetcher, if they are available
	if pixels := p.fetcher.Fetch(currentLine); pixels != nil {
		if p.fetchingSprite {
			p.currentSpritePixelFetched = true
			p.overlayPixels(pixels)
			p.fetchingSprite = false
			p.fetcher.Reset(p.currentFetchPixel, BG_FETCH, nil)
		} else {
			p.shiftInPixels(pixels, currentLine)
			p.fetcher.Reset(p.currentFetchPixel, BG_FETCH, nil)
		}
	}

	if p.canShiftOut() {
		p.shiftOutPixel(currentLine, sprites)
	}
}

func (p *ppu) Reset() {
	p.currentFetchPixel = 0
	p.lcdCurrentPixel = 0
}

func (p *ppu) canShiftOut() bool {
	return len(p.fifo) > 8 && p.lcdCurrentPixel < uint16(SCREEN_WIDTH) && !p.fetchingSprite
}

func (p *ppu) LineFinished() bool {
	return p.lcdCurrentPixel == uint16(SCREEN_WIDTH)
}

func (p *ppu) shiftOutPixel(currentLine int, sprites []SpriteAttribute) {
	if p.isUnfetchedSpritePixel(sprites) {
		return
	} else if p.isUnfetchedWindowPixel(currentLine) {
		// When starting a window fetch clear the entire FIFO and start refetching for window pixels
		p.fifo = make([]RGBPixel, 0)
		p.fetcher.Reset(uint16(p.lcdCurrentPixel), WINDOW_FETCH, nil)
	} else {
		if !p.mmu.BGDisplayEnabled() {
			p.lcdBuffer[currentLine][p.lcdCurrentPixel] = WHITE() // When background is not enabled we should only draw blank pixels
		} else {
			p.lcdBuffer[currentLine][p.lcdCurrentPixel] = p.fifo[0]
		}
		p.fifo = p.fifo[1:]
		p.lcdCurrentPixel += 1
		p.currentSpritePixelFetched = false
	}
}

func (p *ppu) isUnfetchedSpritePixel(sprites []SpriteAttribute) bool {
	for _, sprite := range sprites {
		if sprite != nil && sprite.GetXPosition() > 0 && sprite.GetXPosition()-8 == int(p.lcdCurrentPixel) && !p.currentSpritePixelFetched && p.mmu.SpritesEnabled() {
			p.fetcher.Reset(uint16(p.lcdCurrentPixel), SPRITE_FETCH, sprite)
			p.currentFetchPixel = p.lcdCurrentPixel
			p.fetchingSprite = true
			p.lastFetchedSprite = sprite
			return true
			// 1. Stop current fetching
			// 2. Start fetching sprite pixels
			// 3. Overlay first 8 pixels from background with sprite pixels
			// 4. Resume fetching background
		}
	}
	return false
}

func (p *ppu) isUnfetchedWindowPixel(currentLine int) bool {
	if currentLine >= int(p.mmu.WindowYPosition()) && p.mmu.WindowDisplayEnabled() {
		return true
	} else {
		return false
	}
}

func (p *ppu) shiftInPixels(pixels []RGBPixel, currentLine int) {
	for _, pixel := range pixels {
		p.fifo = append(p.fifo, pixel)
	}
	p.currentFetchPixel += uint16(len(pixels))
}

func (p *ppu) LcdBuffer(y, x int) RGBPixel {
	return p.lcdBuffer[y][x]
}

func (p *ppu) overlayPixels(pixels []RGBPixel) {
	for idx, pixel := range pixels {
		if p.lastFetchedSprite.HasPriority() {
			p.fifo[idx] = p.combineColors(pixel, p.fifo[idx], true)
		} else {
			p.fifo[idx] = p.combineColors(pixel, p.fifo[idx], false)
		}
	}
	p.currentFetchPixel += uint16(len(pixels))
}

func (p *ppu) combineColors(spritePixel, bgPixel RGBPixel, spriteHasPriority bool) RGBPixel {
	// TODO: This might want to check against the value of color 0 rather then WHITE
	if spriteHasPriority || bgPixel == WHITE() {
		return spritePixel
	} else {
		return bgPixel
	}
}
