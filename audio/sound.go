package audio

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sound struct {
	beep rl.Sound
}

func (sound *Sound) Initialize() {
	rl.InitAudioDevice()
	sound.beep = rl.LoadSound("./assets/beep.wav")
}

func (sound *Sound) Close() {
	rl.UnloadSound(sound.beep)
	rl.CloseAudioDevice()
}

func (sound *Sound) Play(provider timerProvider) {
	if provider.GetSoundTimer() > 0 {
		if !rl.IsSoundPlaying(sound.beep) {
			rl.PlaySound(sound.beep)
		}
	} else {
		rl.StopSound(sound.beep)
	}
}
