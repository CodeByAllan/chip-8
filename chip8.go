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
	keyHandler.Initialize()
	sound.Initialize()
	defer sound.Close()

	cpu.Load(*romPath)

	cycleTicker := time.NewTicker(time.Second / time.Duration(cpu.Cycles))
	defer cycleTicker.Stop()

	timerTicker := time.NewTicker(time.Second / time.Duration(cpu.Timer))
	defer timerTicker.Stop()

	for !rl.WindowShouldClose() {
		select {
		case <-cycleTicker.C:
			cpu.Run(keyHandler)

		case <-timerTicker.C:
			cpu.UpdateTimers()
			sound.Play(cpu)

		default:
			display.RenderDisplay()
			keyHandler.HandleInput(cpu)
		}
	}
}
