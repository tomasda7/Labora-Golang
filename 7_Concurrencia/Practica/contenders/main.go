package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var enterKeyPressed bool

	for i := 1; i <= 30 && !enterKeyPressed; i++ {
		fmt.Print("\033[H\033[2J")
		str := "[" + strings.Repeat("=", i) + "🥊" + "]"
		fmt.Println(str)
		time.Sleep(100 * time.Millisecond)
	}

	var input string
	fmt.Scanln(&input)
	enterKeyPressed = true
}
