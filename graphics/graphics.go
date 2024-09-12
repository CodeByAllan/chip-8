package graphics

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	GraphicsWidth  = 64
	GraphicsHeight = 32
	Scale          = 10
)

func RenderGraphics(gfx []byte) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for y := 0; y < GraphicsHeight; y++ {
		for x := 0; x < GraphicsWidth; x++ {
			index := y*GraphicsWidth + x
			color := rl.Black
			if gfx[index] == 1 {
				color = rl.White
			}
			rl.DrawRectangle(
				int32(x*Scale),
				int32(y*Scale),
				int32(Scale),
				int32(Scale),
				color,
			)
		}
	}

	rl.EndDrawing()
}
