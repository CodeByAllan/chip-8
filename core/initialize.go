package core

func (cpu *CPU) Initialize() {
	cpu.PC = 0x200
	cpu.I = 0
	cpu.SoundTimer = 0
	cpu.DelayTimer = 0
	cpu.SP = 0
	cpu.Cycles = 500
	cpu.Timer = 60
	copy(cpu.Mem[:len(fontset)], fontset[:])

}
