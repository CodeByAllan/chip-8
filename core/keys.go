package core

func (cpu *CPU) SetKey(index byte, value byte) {
	if index < 16 {
		cpu.Keys[index] = value
	}
}

func (cpu *CPU) ClearKeys() {
	for i := range cpu.Keys {
		cpu.Keys[i] = 0
	}
}
