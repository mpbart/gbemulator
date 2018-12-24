package main

func GetBit(value uint8, bit uint) int {
	return int((value >> bit) & 1)
}

func GetBitUint16(value uint16, bit uint) int {
	return int((value >> bit) & 1)
}

func BitsToNum(mostSignificantBit, leastSignificantBit int) int {
	if leastSignificantBit == 0 && mostSignificantBit == 0 {
		return 0
	} else if leastSignificantBit == 1 && mostSignificantBit == 0 {
		return 1
	} else if leastSignificantBit == 0 && mostSignificantBit == 1 {
		return 2
	} else {
		return 3
	}
}

func GetHighestBit(num uint8, bits int) int {
	highestBit := 0
	for i := 0; i < bits; i++ {
		if num&0x01 == 1 {
			highestBit = i
		}
		num = num >> 1
	}
	return highestBit
}

func GetHighestInterruptBit(num uint8) int {
	return GetHighestBit(num, 5)
}
