package main

type FetchState int

const (
	TILE_READ FetchState = iota
	READ_DATA_0
	READ_DATA_1
	IDLE
)

type Fetcher interface {
	Fetch() []RGBPixel
}

type fetcher struct {
	currentState           FetchState
	addresser              MemoryAddresser
	mmu                    MMU
	backgroundStartAddress uint16
	windowStartAddress     uint16
	backgroundMapNumber    uint16
	currentTile            uint16
	pixels                 []RGBPixel
}

func createFetcher(mmu MMU) Fetcher {
	return &fetcher{
		currentState: TILE_READ,
		addresser:    CreateMemoryAddresser(mmu.BGAndWindowAddressMode()),
		mmu:          mmu,
		backgroundStartAddress: mmu.BGTileMap(),
		windowStartAddress:     mmu.WindowTileMap(),
		backgroundMapNumber:    0,
		currentTile:            0,
		pixels:                 make([]RGBPixel, 8),
	}
}

func (f *fetcher) Fetch() []RGBPixel {
	switch f.currentState {
	case TILE_READ:
		f.readTile()
	case READ_DATA_0:
		f.readData(0)
	case READ_DATA_1:
		f.readData(1)
	case IDLE:
	}

	if f.shouldChangeState() {
		f.currentState = f.nextState()
		if f.currentState == TILE_READ {
			f.reset()
			pixels := make([]RGBPixel, 8)
			copy(pixels, f.pixels)
			return pixels
		}
	}
	return nil
}

func (f *fetcher) reset() {
	f.backgroundMapNumber = 0
}

func (f *fetcher) nextState() FetchState {
	switch f.currentState {
	case TILE_READ:
		return READ_DATA_0
	case READ_DATA_0:
		return READ_DATA_1
	case READ_DATA_1:
		return IDLE
	case IDLE:
		return TILE_READ
	default:
		return IDLE
	}
}

func (f *fetcher) shouldChangeState() bool {
	return true
}

func (f *fetcher) readTile() {
	f.currentTile = uint16(f.mmu.ReadAt(f.backgroundStartAddress + f.backgroundMapNumber))
}

func (f *fetcher) readData(byteNum int) {
}
