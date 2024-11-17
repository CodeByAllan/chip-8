package core

import "chip-8/graphics"

type CPU struct {
	V          [16]byte
	I          uint16
	PC         uint16
	Mem        [4096]byte
	Display    *graphics.Display
	DelayTimer byte
	SoundTimer byte
	Stack      [16]uint16
	SP         uint8
	Keys       [16]byte
	Cycles     int
	Timer      int
}
