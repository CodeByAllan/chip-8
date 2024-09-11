package chip8

import (
	"chip-8/common"
	"chip-8/cpu"
)

func Chip8() {
	cpuInstance := &common.CPU{}
	cpu.Initialize(cpuInstance)
}
