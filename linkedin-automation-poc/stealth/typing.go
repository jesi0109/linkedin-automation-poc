package stealth

import (
	"fmt"
	"time"
)

func TypeHumanLike(text string) {
	for _, c := range text {
		fmt.Print(string(c))
		time.Sleep(120 * time.Millisecond)
	}
	fmt.Println()
}
