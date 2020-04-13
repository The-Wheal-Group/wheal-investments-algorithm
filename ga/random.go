package ga

import (
	"math/rand"
	"time"
)

func Random() *rand.Rand {
	randomSource := rand.NewSource(time.Now().UnixNano())
	return rand.New(randomSource)
}
