package graphics

import rl "github.com/gen2brain/raylib-go/raylib"

type Display struct {
	DisplayWidth    int
	DisplayHeight   int
	Screen          [64 * 32]byte
	scale           int
	backgroundColor rl.Color
	spriteColor     rl.Color
}
