package core

func (cpu *CPU) UpdateTimers() {
	if cpu.DelayTimer > 0 {
		cpu.DelayTimer--
	}
	if cpu.SoundTimer > 0 {
		cpu.SoundTimer--
	}
}
