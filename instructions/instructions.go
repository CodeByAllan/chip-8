package instructions

import "chip-8/common"

func ClearDisplay(cpu *common.CPU) {
	for i := range cpu.Screen {
		cpu.Screen[i] = 0
	}
}
func ReturnFromSubroutine(cpu *common.CPU) {
	if cpu.SP == 0 {
		return
	}
	cpu.SP--
	cpu.PC = cpu.Stack[cpu.SP]
}
func JumpAddress(cpu *common.CPU, opcode uint16) {
	cpu.PC = opcode & 0x0FFF
}
func SubRoutine(cpu *common.CPU, opcode uint16) {
	cpu.Stack[cpu.SP] = cpu.PC
	cpu.SP++
	JumpAddress(cpu, opcode)

}
func SkipIfEqual(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	if cpu.V[x] == byte(nn) {
		cpu.PC += 2
	}
}
func SkipIfNotEqual(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	if cpu.V[x] != byte(nn) {
		cpu.PC += 2
	}
}
func SkipIfRegistersEqual(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	if cpu.V[x] == cpu.V[y] {
		cpu.PC += 2
	}
}
func LoadImmediate(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	cpu.V[x] = byte(nn)
}

func AddImmediate(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	cpu.V[x] += byte(nn)

}
