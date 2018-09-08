package main

type AddressMode int

const (
	ADDRESS_MODE_8000 AddressMode = iota
	ADDRESS_MODE_8800
)

type MemoryAddresser interface {
	GetAddress(uint8) uint16
}

type memoryAddresser struct {
	addressMode AddressMode
}

func CreateMemoryAddresser(addressMode AddressMode) MemoryAddresser {
	return &memoryAddresser{
		addressMode: addressMode,
	}
}

func (a *memoryAddresser) GetAddress(offset uint8) uint16 {
	if a.addressMode == ADDRESS_MODE_8000 {
		return uint16(0x8000 + uint16(offset))
	} else {
		if offset < 128 {
			return uint16(0x9000 + uint16(offset))
		} else {
			return uint16(0x8800 + uint16(offset))
		}
	}
}
