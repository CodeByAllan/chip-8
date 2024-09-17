package core

import (
	"chip-8/keyboard"
	"chip-8/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func clearDisplay(cpu *CPU) {
	for i := range cpu.Display.Screen {
		cpu.Display.Screen[i] = 0
	}
}
func returnFromSubroutine(cpu *CPU) {
	if cpu.SP == 0 {
		return
	}
	cpu.SP--
	cpu.PC = cpu.Stack[cpu.SP]
}
func jumpAddress(cpu *CPU, opcode uint16) {
	cpu.PC = opcode & 0x0FFF
}
func subRoutine(cpu *CPU, opcode uint16) {
	cpu.Stack[cpu.SP] = cpu.PC
	cpu.SP++
	jumpAddress(cpu, opcode)

}
func skipIfEqual(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	if cpu.V[x] == byte(nn) {
		cpu.PC += 2
	}
}
func skipIfNotEqual(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	if cpu.V[x] != byte(nn) {
		cpu.PC += 2
	}
}
func skipIfRegistersEqual(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	if cpu.V[x] == cpu.V[y] {
		cpu.PC += 2
	}
}
func loadImmediate(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	cpu.V[x] = byte(nn)
}
func addImmediate(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	nn := opcode & 0x00FF
	cpu.V[x] += byte(nn)

}
func assignRegisterValue(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[y]
}
func orRegisterValue(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[x] | cpu.V[y]
}
func andRegisterValue(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[x] & cpu.V[y]
}
func xorRegisterValue(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	cpu.V[x] = cpu.V[x] ^ cpu.V[y]
}
func addRegisterWithCarry(cpu *CPU, opcode uint16) {
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
func subtractRegisterValue(cpu *CPU, opcode uint16) {
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
func shiftRightRegisterValue(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.V[0xF] = cpu.V[x] & 0x01
	cpu.V[x] >>= 1
}
func subtractRegisterFromValue(cpu *CPU, opcode uint16) {
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
func shiftLeftRegisterValue(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.V[0xF] = (cpu.V[x] & 0x80) >> 7
	cpu.V[x] <<= 1
}
func skipIfRegistersNotEqual(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	y := (opcode & 0x00F0) >> 4
	if cpu.V[x] != cpu.V[y] {
		cpu.PC += 2
	}
}
func assignIndexRegister(cpu *CPU, opcode uint16) {
	cpu.I = opcode & 0x0FFF
}
func skipToAddressWithOffset(cpu *CPU, opcode uint16) {
	offset := opcode & 0x0FFF
	cpu.PC = offset + uint16(cpu.V[0])
}
func setRegisterIfRandomEquals(cpu *CPU, opcode uint16) {
	i := (opcode & 0x0F00) >> 8
	x := utils.GenerateRandom8Bit()
	nn := opcode & 0x00FF
	if x == uint8(nn) {
		cpu.V[i] = 1
	} else {
		cpu.V[i] = 0
	}
}
func updateTimers(cpu *CPU) {
	if cpu.DelayTimer > 0 {
		cpu.DelayTimer--
	}
	if cpu.SoundTimer > 0 {
		cpu.SoundTimer--
	}
}
func drawSprite(cpu *CPU, opcode uint16) {
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

				if cpu.Display.Screen[yPos*64+xPos] == 1 {
					cpu.V[0xF] = 1
				}

				cpu.Display.Screen[yPos*64+xPos] ^= 1
			}
		}
	}
}
func skipIfKeyPressed(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	if cpu.Keys[cpu.V[x]] == 1 {
		cpu.PC += 2
	}
}
func ignoreIfKeyPressed(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	if cpu.Keys[cpu.V[x]] != 1 {
		cpu.PC += 2
	}
}
func setVXFromDelayTimer(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.V[x] = cpu.DelayTimer
}
func waitForKeyPressAndStoreInVX(cpu *CPU, opcode uint16, handler *keyboard.Handler) {
	x := (opcode & 0x0F00) >> 8

	for !handler.AnyKeyPressed() {
		handler.HandleInput(cpu)
	}
	for key, value := range handler.KeyMap {
		if rl.IsKeyDown(key) {
			cpu.V[x] = value
			break
		}
	}
}
func setDelayTimerFromVX(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.DelayTimer = cpu.V[x]
}
func setSoundTimerFromVX(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.SoundTimer = cpu.V[x]
}
func setVXToI(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	cpu.I += uint16(cpu.V[x])
}
func setIToSpriteLocation(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	digit := cpu.V[x]
	cpu.I = uint16(digit) * 5
}
func storeBCD(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	value := cpu.V[x]
	cpu.Mem[cpu.I] = value / 100
	cpu.Mem[cpu.I+1] = (value / 10) % 10
	cpu.Mem[cpu.I+2] = value % 10
}
func storeRegisters(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	for i := uint16(0); i <= x; i++ {
		cpu.Mem[cpu.I+i] = cpu.V[i]
	}
}
func loadRegisters(cpu *CPU, opcode uint16) {
	x := (opcode & 0x0F00) >> 8
	for i := uint16(0); i <= x; i++ {
		cpu.V[i] = cpu.Mem[cpu.I+i]
	}
}
