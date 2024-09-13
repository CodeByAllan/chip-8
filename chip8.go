package chip8

import (
	"chip-8/audio"
	"chip-8/common"
	"chip-8/cpu"
	"chip-8/graphics"
	"chip-8/keyboard"
	"chip-8/rom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Chip8(romPath *string) {
	rl.InitWindow(graphics.GraphicsWidth*graphics.Scale, graphics.GraphicsHeight*graphics.Scale, "CHIP-8 Emulator")
	audio.InitAudio()
	defer audio.CloseAudio()
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	cpuInstance := &common.CPU{}
	keyHandler := &keyboard.Handler{}
	cpu.Initialize(cpuInstance)
	rom.Load(cpuInstance, *romPath)
	keyHandler.Initialize()
	for !rl.WindowShouldClose() {
		cpu.Run(cpuInstance)
		audio.Audio(cpuInstance)
		graphics.RenderGraphics(cpuInstance.Screen[:])
		keyHandler.HandleInput(cpuInstance)
	}

}
