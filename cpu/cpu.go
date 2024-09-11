package cpu

import (
	"chip-8/fontset"
	"chip-8/instructions"
	"fmt"
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
func Run(cpu *CPU) {
	for {
		opcode := uint16(cpu.Mem[cpu.PC])<<8 | uint16(cpu.Mem[cpu.PC+1])
		fmt.Printf("Opcode: 0x%X\n", opcode)
		cpu.PC += 2
		switch opcode & 0xF000 {
		case 0x0000:
			switch opcode & 0x00FF {
			case 0x0E0:
				instructions.ClearDisplay(&cpu.Screen)
			default:
				fmt.Printf("Opcode desconhecido: 0x%X\n", opcode)
			}
		default:
			fmt.Printf("Opcode desconhecido: 0x%X\n", opcode)
		}
	}
}
