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
	tileData               uint16
	doAction               bool
	pixels                 []RGBPixel
}

func createFetcher(mmu MMU) Fetcher {
	return &fetcher{
		currentState:           TILE_READ,
		addresser:              CreateMemoryAddresser(mmu.BGAndWindowAddressMode()),
		mmu:                    mmu,
		backgroundStartAddress: mmu.BGTileMap(),
		windowStartAddress:     mmu.WindowTileMap(),
		backgroundMapNumber:    0,
		currentTile:            0,
		tileData:               0,
		doAction:               false,
		pixels:                 make([]RGBPixel, 8),
	}
}

// TODO: Add checking for and writing of pixels and windows
func (f *fetcher) Fetch() []RGBPixel {
	if !f.canRun() {
		return nil
	}

	switch f.currentState {
	case TILE_READ:
		f.readTile()
	case READ_DATA_0:
		f.readData(0)
	case READ_DATA_1:
		f.readData(1)
	case IDLE:
		f.setPixels()
	}

	f.currentState = f.nextState()

	if f.currentState == TILE_READ {
		pixels := make([]RGBPixel, 8)
		copy(pixels, f.pixels)
		f.reset()
		return pixels
	}
	return nil
}

func (f *fetcher) reset() {
	f.backgroundMapNumber = 0
	for i := 0; i < len(f.pixels); i++ {
		f.pixels[i] = WHITE()
	}
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

func (f *fetcher) readTile() {
	f.currentTile = uint16(f.mmu.ReadAt(f.backgroundStartAddress + f.backgroundMapNumber))
	// TODO: I don't know if this is right...
	f.backgroundMapNumber += 1
}

func (f *fetcher) readData(byteNum uint8) {
	f.tileData += uint16(f.mmu.ReadAt(f.addresser.GetAddress(uint8(f.currentTile)+byteNum))) << (8 * byteNum)
}

func (f *fetcher) setPixels() {
	for i := 0; i < len(f.pixels); i++ {
		f.pixels[i] = f.getColor(i)
	}
}

func (f *fetcher) getColor(i int) RGBPixel {
	lowerBit := GetBitUint16(f.tileData, uint(i))
	upperBit := GetBitUint16(f.tileData, uint(i+7))
	return f.mmu.ConvertNumToPixel(BitsToNum(upperBit, lowerBit))
}

// Method to run fetcher at half speed
func (f *fetcher) canRun() bool {
	oldValue := f.doAction
	f.doAction = !f.doAction
	return oldValue
}
