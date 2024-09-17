package core

import (
	"chip-8/keyboard"
	"fmt"
)

func (cpu *CPU) fetch() uint16 {
	opcode := uint16(cpu.Mem[cpu.PC])<<8 | uint16(cpu.Mem[cpu.PC+1])
	cpu.PC += 2
	return opcode
}
func (cpu *CPU) system(opcode uint16) {
	switch opcode & 0x00FF {
	case 0x00E0:
		clearDisplay(cpu)
	case 0x00EE:
		returnFromSubroutine(cpu)
	default:
		fmt.Printf("Unknown opcode: 0x%X\n", opcode)
	}
}
func (cpu *CPU) flowControl(opcode uint16) {
	switch opcode & 0xF000 {
	case 0x1000:
		jumpAddress(cpu, opcode)
	case 0x2000:
		subRoutine(cpu, opcode)
	case 0x3000:
		skipIfEqual(cpu, opcode)
	case 0x4000:
		skipIfNotEqual(cpu, opcode)
	case 0x5000:
		skipIfRegistersEqual(cpu, opcode)
	case 0x9000:
		skipIfRegistersNotEqual(cpu, opcode)
	default:
		fmt.Printf("Unknown flow control opcode: 0x%X\n", opcode)
	}
}
func (cpu *CPU) register(opcode uint16) {
	switch opcode & 0xF000 {
	case 0x6000:
		loadImmediate(cpu, opcode)
	case 0x7000:
		addImmediate(cpu, opcode)
	default:
		fmt.Printf("Unknown register opcode: 0x%X\n", opcode)
	}
}
func (cpu *CPU) mathLogic(opcode uint16) {
	switch opcode & 0x000F {
	case 0x0000:
		assignRegisterValue(cpu, opcode)
	case 0x0001:
		orRegisterValue(cpu, opcode)
	case 0x0002:
		andRegisterValue(cpu, opcode)
	case 0x0003:
		xorRegisterValue(cpu, opcode)
	case 0x0004:
		addRegisterWithCarry(cpu, opcode)
	case 0x0005:
		subtractRegisterValue(cpu, opcode)
	case 0x0006:
		shiftRightRegisterValue(cpu, opcode)
	case 0x0007:
		subtractRegisterFromValue(cpu, opcode)
	case 0x000E:
		shiftLeftRegisterValue(cpu, opcode)
	default:
		fmt.Printf("Unknown opcode: 0x%X\n", opcode)
	}
}
func (cpu *CPU) memory(opcode uint16) {
	switch opcode & 0xF000 {
	case 0xA000:
		assignIndexRegister(cpu, opcode)
	case 0xB000:
		skipToAddressWithOffset(cpu, opcode)
	default:
		fmt.Printf("Unknown memory opcode: 0x%X\n", opcode)
	}
}
func (cpu *CPU) random(opcode uint16) {
	setRegisterIfRandomEquals(cpu, opcode)
}
func (cpu *CPU) drawing(opcode uint16) {
	drawSprite(cpu, opcode)
}
func (cpu *CPU) keyboard(opcode uint16) {
	switch opcode & 0x00FF {
	case 0x009E:
		skipIfKeyPressed(cpu, opcode)
	case 0x00A1:
		ignoreIfKeyPressed(cpu, opcode)
	default:
		fmt.Printf("Unknown keyboard opcode: 0x%X\n", opcode)
	}
}
func (cpu *CPU) timerAndMemory(opcode uint16, keyhandler *keyboard.Handler) {
	switch opcode & 0x00FF {
	case 0x0007:
		setVXFromDelayTimer(cpu, opcode)
	case 0x000A:
		waitForKeyPressAndStoreInVX(cpu, opcode, keyhandler)
	case 0x0015:
		setDelayTimerFromVX(cpu, opcode)
	case 0x0018:
		setSoundTimerFromVX(cpu, opcode)
	case 0x001E:
		setVXToI(cpu, opcode)
	case 0x0029:
		setIToSpriteLocation(cpu, opcode)
	case 0x0033:
		storeBCD(cpu, opcode)
	case 0x0055:
		storeRegisters(cpu, opcode)
	case 0x0065:
		loadRegisters(cpu, opcode)
	default:
		fmt.Printf("Unknown timer/memory opcode: 0x%X\n", opcode)
	}
}
