package main

type FetchState int
type FetchMode int

const (
	TILE_READ FetchState = iota
	READ_DATA_0
	READ_DATA_1
	IDLE

	BG_FETCH FetchMode = iota
	SPRITE_FETCH
	WINDOW_FETCH
)

type Fetcher interface {
	Fetch(int) []RGBPixel
	Reset(uint16, FetchMode)
}

type fetcher struct {
	currentState           FetchState
	addresser              MemoryAddresser
	mmu                    MMU
	fetchMode              FetchMode
	backgroundStartAddress uint16
	windowStartAddress     uint16
	currentTile            uint16
	currentPixel           uint16
	tileData               uint16
	doAction               bool
	pixels                 []RGBPixel
}

func createFetcher(mmu MMU) Fetcher {
	return &fetcher{
		currentState:           TILE_READ,
		addresser:              CreateMemoryAddresser(mmu.BGAndWindowAddressMode()),
		mmu:                    mmu,
		fetchMode:              BG_FETCH,
		backgroundStartAddress: mmu.BGTileMap(),
		windowStartAddress:     mmu.WindowTileMap(),
		currentPixel:           0,
		currentTile:            0,
		tileData:               0,
		doAction:               false,
		pixels:                 make([]RGBPixel, 8),
	}
}

func (f *fetcher) Fetch(currentLine int) []RGBPixel {
	if !f.canRun() {
		return nil
	}

	switch f.currentState {
	case TILE_READ:
		f.readTile(currentLine)
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
		f.Reset(f.currentPixel+8, BG_FETCH) // May need to do something for fetch mode if fetching for window
		return pixels
	}
	return nil
}

// Have reset take a param so that the ppu can tell it where to start fetching sprite pixels at
func (f *fetcher) Reset(currentPixel uint16, fetchMode FetchMode) {
	f.currentPixel = currentPixel
	f.fetchMode = fetchMode
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

func (f *fetcher) readTile(currentLine int) {
	if f.fetchMode == BG_FETCH {
		yOffset, xOffset := uint16(currentLine>>3), uint16(f.currentPixel>>3)
		f.currentTile = uint16(f.mmu.ReadAt(f.backgroundStartAddress + yOffset + xOffset))
	} else if f.fetchMode == SPRITE_FETCH {
	}
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

// Use this to index based on fetching bg vs. sprite vs. window
func (f *fetcher) tileBaseAddress() uint16 {
	return 0
}
