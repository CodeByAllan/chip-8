package cpu

import (
	"chip-8/fontset"
)

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

func Initialize(cpu *CPU) {
	cpu.PC = 0x200
	cpu.I = 0
	cpu.SoundTimer = 0
	cpu.DelayTimer = 0
	cpu.SP = 0
	copy(cpu.Mem[:len(fontset.Chip8Fontset)], fontset.Chip8Fontset[:])

}
