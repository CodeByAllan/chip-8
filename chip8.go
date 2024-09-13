package chip8

import (
	"chip-8/audio"
	"chip-8/common"
	"chip-8/cpu"
	"chip-8/graphics"
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
	cpu.Initialize(cpuInstance)
	rom.Load(cpuInstance, *romPath)

	stop := make(chan struct{})
	go func() {
		cpu.Run(cpuInstance, stop)
	}()

	for !rl.WindowShouldClose() {
		audio.Audio(cpuInstance)
		graphics.RenderGraphics(cpuInstance.Screen[:])
	}

	close(stop)
}
