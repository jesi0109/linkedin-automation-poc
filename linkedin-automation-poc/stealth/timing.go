package stealth

import (
	"math/rand"
	"time"
)

func HumanDelay() {
	delay := rand.Intn(2000) + 800
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
