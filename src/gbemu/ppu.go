package main

type PPU interface {
	Tick([]SpriteAttribute)
	LineFinished() bool
}

type ppu struct {
	fifo    []uint8
	fetcher Fetcher
}

func createPPU(mmu MMU) PPU {
	return &ppu{
		fifo:    make([]uint8, 16),
		fetcher: createFetcher(mmu),
	}
}

func (p *ppu) Tick(_ []SpriteAttribute) {
	p.fetcher.Fetch()

	if p.canShift() {
		p.shiftOutPixel()
	}
}

func (p *ppu) canShift() bool {
	return len(p.fifo) > 8
}

func (p *ppu) LineFinished() bool {
	return false
}

func (p *ppu) shiftOutPixel() {
}
