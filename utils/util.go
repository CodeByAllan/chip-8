package utils

import (
	"math/rand"
	"time"
)

func newRNG() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateRandom8Bit() uint8 {
	rng := newRNG()
	return uint8(rng.Intn(256))
}
