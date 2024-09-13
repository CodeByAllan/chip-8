package instructions

import (
	"chip-8/common"
	"chip-8/utils"
)

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
func AssignRegisterValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[y]
}
func OrRegisterValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[x] | cpu.V[y]
}
func AndRegisterValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[x] & cpu.V[y]
}
func XorRegisterValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[x] ^ cpu.V[y]
}
func AddRegisterWithCarry(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	sum := uint16(cpu.V[x]) + uint16(cpu.V[y])
	cpu.V[x] = byte(sum)
	if sum > 0xFF {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}
}
func SubtractRegisterValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	vx := uint16(cpu.V[x])
	vy := uint16(cpu.V[y])
	result := vx - vy
	cpu.V[x] = byte(result)
	if vy > vx {
		cpu.V[0xF] = 0
	} else {
		cpu.V[0xF] = 1
	}
}
func ShiftRightRegisterValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.V[0xF] = cpu.V[x] & 0x01
	cpu.V[x] >>= 1
}
func SubtractRegisterFromValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	vx := uint16(cpu.V[x])
	vy := uint16(cpu.V[y])

	result := vy - vx
	cpu.V[x] = byte(result)

	if vy > vx {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}
}
func ShiftLeftRegisterValue(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.V[0xF] = (cpu.V[x] & 0x80) >> 7
	cpu.V[x] <<= 1
}
func SkipIfRegistersNotEqual(cpu *common.CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	if cpu.V[x] != cpu.V[y] {
		cpu.PC += 2
	}
}
func AssignIndexRegister(cpu *common.CPU, opcode uint16) {
	cpu.I = opcode & 0x0FFF
}
func SkipToAddressWithOffset(cpu *common.CPU, opcode uint16) {
	offset := opcode & 0x0FFF
	cpu.PC = offset + uint16(cpu.V[0])
}
func SetRegisterIfRandomEquals(cpu *common.CPU, opcode uint16) {
	i := (opcode & 0x0F00) >> 8
	x := utils.GenerateRandom8Bit()
	nn := opcode & 0x00FF
	if x == uint8(nn) {
		cpu.V[i] = 1
	} else {
		cpu.V[i] = 0
	}
}
func UpdateTimers(cpu *common.CPU) {
	if cpu.DelayTimer > 0 {
		cpu.DelayTimer--
	}
	if cpu.SoundTimer > 0 {
		cpu.SoundTimer--
	}
}
func DrawSprite(cpu *common.CPU, opcode uint16) {
	x := cpu.V[(opcode&0x0F00)>>8]
	y := cpu.V[(opcode&0x00F0)>>4]
	n := opcode & 0x000F

	cpu.V[0xF] = 0

	for row := uint16(0); row < uint16(n); row++ {
		spriteLine := cpu.Mem[cpu.I+row]

		for col := uint16(0); col < 8; col++ {
			if (spriteLine & (0x80 >> col)) != 0 {
				xPos := (int(x) + int(col)) % 64
				yPos := (int(y) + int(row)) % 32

				if xPos < 0 {
					xPos += 64
				}
				if yPos < 0 {
					yPos += 32
				}

				if cpu.Screen[yPos*64+xPos] == 1 {
					cpu.V[0xF] = 1
				}

				cpu.Screen[yPos*64+xPos] ^= 1
			}
		}
	}
}
