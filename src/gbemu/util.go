package main

func GetBit(value uint8, bit uint) int {
	return int((value >> bit) & 1)
}
