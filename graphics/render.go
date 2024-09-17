package graphics

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (display *Display) clearScreen() {
	rl.ClearBackground(display.backgroundColor)
}

func (display *Display) drawPixel(x, y int, color rl.Color) {
	rl.DrawRectangle(
		int32(x*display.scale),
		int32(y*display.scale),
		int32(display.scale),
		int32(display.scale),
		color,
	)
}
func (display *Display) RenderDisplay() {
	rl.BeginDrawing()
	display.clearScreen()
	for y := 0; y < display.DisplayHeight; y++ {
		for x := 0; x < display.DisplayWidth; x++ {
			index := y*display.DisplayWidth + x
			color := display.backgroundColor
			if display.Screen[index] == 1 {
				color = display.spriteColor
			}
			display.drawPixel(x, y, color)
		}
	}

	rl.EndDrawing()
}
