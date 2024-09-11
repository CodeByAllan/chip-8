package cpu

import (
	"chip-8/common"
	"chip-8/fontset"
	"chip-8/instructions"
	"fmt"
)

func Initialize(cpu *common.CPU) {
	cpu.PC = 0x200
	cpu.I = 0
	cpu.SoundTimer = 0
	cpu.DelayTimer = 0
	cpu.SP = 0
	copy(cpu.Mem[:len(fontset.Chip8Fontset)], fontset.Chip8Fontset[:])

}
func Run(cpu *common.CPU) {
	for {
		opcode := uint16(cpu.Mem[cpu.PC])<<8 | uint16(cpu.Mem[cpu.PC+1])
		fmt.Printf("Opcode: 0x%X\n", opcode)
		cpu.PC += 2
		switch opcode & 0xF000 {
		case 0x0000:
			switch opcode & 0x00FF {
			case 0x0E0:
				instructions.ClearDisplay(cpu)
			case 0x0EE:
				instructions.ReturnFromSubroutine(cpu)
			default:
				fmt.Printf("Opcode desconhecido: 0x%X\n", opcode)
			}
		default:
			fmt.Printf("Opcode desconhecido: 0x%X\n", opcode)
		}
	}
}
