package chip8

import "chip-8/cpu"

func Chip8() {
	cpuInstance := &cpu.CPU{}
	cpu.Initialize(cpuInstance)
}
