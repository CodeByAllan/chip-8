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
		case 0x1000:
			instructions.JumpAddress(cpu, opcode)
		case 0x2000:
			instructions.SubRoutine(cpu, opcode)
		case 0x3000:
			instructions.SkipIfEqual(cpu, opcode)
		case 0x4000:
			instructions.SkipIfNotEqual(cpu, opcode)
		case 0x5000:
			instructions.SkipIfRegistersEqual(cpu, opcode)
		case 0x6000:
			instructions.LoadImmediate(cpu, opcode)
		case 0x7000:
			instructions.AddImmediate(cpu, opcode)
		case 0x8000:
			switch opcode & 0x000F {
			case 0x0000:
				instructions.AssignRegisterValue(cpu, opcode)
			case 0x0001:
				instructions.OrRegisterValue(cpu, opcode)
			case 0x0002:
				instructions.AndRegisterValue(cpu, opcode)
			case 0x0003:
				instructions.XorRegisterValue(cpu, opcode)
			case 0x0004:
				instructions.AddRegisterWithCarry(cpu, opcode)
			case 0x0005:
				instructions.SubtractRegisterValue(cpu, opcode)
			case 0x0006:
				instructions.ShiftRightRegisterValue(cpu, opcode)
			case 0x0007:
				instructions.SubtractRegisterFromValue(cpu, opcode)
			case 0x000E:
				instructions.ShiftLeftRegisterValue(cpu, opcode)
			default:
				fmt.Printf("Opcode desconhecido: 0x%X\n", opcode)
			}
		case 0x9000:
			instructions.SkipIfRegistersNotEqual(cpu, opcode)
		case 0xA000:
			instructions.AssignIndexRegister(cpu, opcode)
		case 0xB000:
			instructions.SkipToAddressWithOffset(cpu, opcode)
		default:
			fmt.Printf("Opcode desconhecido: 0x%X\n", opcode)
		}
	}
}
