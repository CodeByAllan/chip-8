package graphics

import rl "github.com/gen2brain/raylib-go/raylib"

func (display *Display) Initialize() {
	display.DisplayWidth = 64
	display.DisplayHeight = 32
	display.scale = 10
	display.backgroundColor = rl.Black
	display.spriteColor = rl.White
}
