package main

import "fmt"

type SpriteAttribute interface {
	GetYPosition() int
	GetXPosition() int
	GetTileNumber() int
	HasPriority() bool
	HorizontalFlip() bool
	VerticalFlip() bool
	PaletteNumber() int
}

type spriteAttribute struct {
	yPosition  uint8
	xPosition  uint8
	tileNumber uint8
	flags      uint8
}

func fromBytes(bytes []uint8) SpriteAttribute {
	if len(bytes) != 4 {
		fmt.Printf("ERROR: Attempting to init sprite attribute with wrong number of bytes (%v)\n", len(bytes))
		return &spriteAttribute{}
	}
	return &spriteAttribute{
		yPosition:  bytes[0],
		xPosition:  bytes[1],
		tileNumber: bytes[2],
		flags:      bytes[3],
	}
}

func (s *spriteAttribute) GetYPosition() int {
	return int(s.yPosition)
}

func (s *spriteAttribute) GetXPosition() int {
	return int(s.xPosition)
}

func (s *spriteAttribute) GetTileNumber() int {
	return int(s.tileNumber)
}

func (s *spriteAttribute) HasPriority() bool {
	return s.flags&0x80 == 0
}

func (s *spriteAttribute) HorizontalFlip() bool {
	return s.flags&0x20 == 0x20
}

func (s *spriteAttribute) VerticalFlip() bool {
	return s.flags&0x40 == 0x40
}

func (s *spriteAttribute) PaletteNumber() int {
	if s.flags&0x10 == 0x10 {
		return 1
	} else {
		return 0
	}
}
