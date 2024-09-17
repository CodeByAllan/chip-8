package core

import (
	"chip-8/keyboard"
	"fmt"
)

func (cpu *CPU) Run(keyhandler *keyboard.Handler) {
	opcode := cpu.fetch()
	cpu.execute(opcode, keyhandler)
}

func (cpu *CPU) execute(opcode uint16, keyhandler *keyboard.Handler) {
	switch opcode & 0xF000 {
	case 0x0000:
		cpu.system(opcode)
	case 0x1000, 0x2000, 0x3000, 0x4000, 0x5000, 0x9000:
		cpu.flowControl(opcode)
	case 0x6000, 0x7000:
		cpu.register(opcode)
	case 0x8000:
		cpu.mathLogic(opcode)
	case 0xA000, 0xB000:
		cpu.memory(opcode)
	case 0xC000:
		cpu.random(opcode)
	case 0xD000:
		cpu.drawing(opcode)
	case 0xE000:
		cpu.keyboard(opcode)
	case 0xF000:
		cpu.timerAndMemory(opcode, keyhandler)
	default:
		fmt.Printf("Unknown opcode: 0x%X\n", opcode)
	}
}
