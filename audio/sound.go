package audio

import (
	_ "embed" // Import necessário para usar go:embed

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed ..\assets\beep.wav
var beepWav []byte

type Sound struct {
	beep rl.Sound
}

func (sound *Sound) Initialize() {
	rl.InitAudioDevice()
	wavData := rl.LoadWaveFromMemory(".wav", beepWav, int32(len(beepWav)))
	sound.beep = rl.LoadSoundFromWave(wavData)
	rl.UnloadWave(wavData)
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
