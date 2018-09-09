package main

func GetBit(value uint8, bit uint) int {
	return int((value >> bit) & 1)
}

func GetBitUint16(value uint16, bit uint) int {
	return int((value >> bit) & 1)
}

func BitsToNum(highBit, lowBit int) int {
	if lowBit == 0 && highBit == 0 {
		return 0
	} else if lowBit == 1 && highBit == 0 {
		return 1
	} else if lowBit == 0 && highBit == 1 {
		return 2
	} else {
		return 3
	}
}
