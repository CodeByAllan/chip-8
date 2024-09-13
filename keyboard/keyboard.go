package keyboard

import (
	"chip-8/common"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Handler struct {
	keyMap map[int32]byte
}

func (handler *Handler) Initialize() {
	handler.keyMap = map[int32]byte{
		rl.KeyOne: 0x1, rl.KeyTwo: 0x2, rl.KeyThree: 0x3, rl.KeyC: 0xC,
		rl.KeyQ: 0x4, rl.KeyW: 0x5, rl.KeyE: 0x6, rl.KeyR: 0xD,
		rl.KeyA: 0x7, rl.KeyS: 0x8, rl.KeyD: 0x9, rl.KeyF: 0xE,
		rl.KeyZ: 0xA, rl.KeyX: 0x0, rl.KeyV: 0xF,
	}

}
func (handler *Handler) HandleInput(cpu *common.CPU) {
	for i := 0; i < 16; i++ {
		cpu.Keys[i] = 0
	}
	for key, value := range handler.keyMap {
		if rl.IsKeyDown(key) {
			cpu.Keys[value] = 1
		}
	}
}
