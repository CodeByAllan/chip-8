package chip8

import (
	"chip-8/audio"
	"chip-8/core"
	"chip-8/graphics"
	"chip-8/keyboard"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Run(romPath *string) {
	rl.InitWindow(64*10, 32*10, "CHIP-8 Emulator")

	defer rl.CloseWindow()
	display := &graphics.Display{}
	cpu := &core.CPU{Display: display}
	keyHandler := &keyboard.Handler{}
	sound := &audio.Sound{}

	cpu.Initialize()
	display.Initialize()
	sound.Initialize()

	defer sound.Close()

	cpu.Load(*romPath)

	keyHandler.Initialize()

	lastTimerUpdate := time.Now()

	for !rl.WindowShouldClose() {
		cpu.Run(keyHandler)
		cpu.UpdateTimersIfNeeded(&lastTimerUpdate)
		sound.Play(cpu)
		display.RenderDisplay()
		keyHandler.HandleInput(cpu)
	}
}
