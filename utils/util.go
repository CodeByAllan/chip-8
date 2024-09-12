package utils

import (
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandom8Bit() uint8 {
	return uint8(rng.Intn(256))
}
