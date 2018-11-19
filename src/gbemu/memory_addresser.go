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
	mmu MMU
}

func CreateMemoryAddresser(mmu MMU) MemoryAddresser {
	return &memoryAddresser{
		mmu: mmu,
	}
}

func (a *memoryAddresser) GetAddress(offset uint8) uint16 {
	if a.getAddressMode() == ADDRESS_MODE_8000 {
		return uint16(0x8000 + uint16(offset)<<4)
	} else {
		if offset < 128 {
			return uint16(0x9000 + uint16(offset)<<4)
		} else {
			return uint16(0x8800 + uint16(offset)<<4)
		}
	}
}

func (a *memoryAddresser) getAddressMode() AddressMode {
	return a.mmu.BGAndWindowAddressMode()
}
