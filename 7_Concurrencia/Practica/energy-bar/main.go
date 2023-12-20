package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/eiannone/keyboard" //https://github.com/eiannone/keyboard
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	stopChan := make(chan bool)

	go func() {
		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if key == keyboard.KeyEnter {
				stopChan <- true
				return
			}
		}
	}()

	printPunch(stopChan)
}

func printPunch(stopChan <-chan bool) {
	limit := 30
	for i := 1; i <= limit*2+1; i++ {
		// select: https://go.dev/tour/concurrency/5
		select {
		case <-stopChan:
			fmt.Println("Power:", i)
			return
		default:
			fmt.Print("\033[H\033[2J")
			str := "[" + strings.Repeat("=", i%(limit+1)) + "🥊" + "]"
			fmt.Println(str)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

