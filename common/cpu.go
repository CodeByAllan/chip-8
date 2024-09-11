package common

type CPU struct {
	V          [16]byte
	I          uint16
	PC         uint16
	Mem        [4096]byte
	Screen     [64 * 32]byte
	DelayTimer byte
	SoundTimer byte
	Stack      [16]uint16
	SP         uint8
}
