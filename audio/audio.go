package audio

import (
	"chip-8/common"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var beep rl.Sound

func InitAudio() {
	rl.InitAudioDevice()
	beep = rl.LoadSound("./assets/beep.wav")
}

func CloseAudio() {
	rl.UnloadSound(beep)
	rl.CloseAudioDevice()
}

func Audio(cpu *common.CPU) {
	if cpu.SoundTimer > 0 {
		if !rl.IsSoundPlaying(beep) {
			rl.PlaySound(beep)
		}
	} else {
		rl.StopSound(beep)
	}
}
