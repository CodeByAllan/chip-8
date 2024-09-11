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
