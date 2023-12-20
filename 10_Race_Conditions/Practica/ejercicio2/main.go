package main

import (
	"fmt"
	"sync"
)

/*
Lo que quiero que hagan es un programita que use canales para enviar desde un hilo los caracteres en el orden inverso,
usando la función defer, y en el otro hilo ir recibiendolos para terminar armando el string invertido.
*/

func main() {
	var wg sync.WaitGroup
	var runesChannel chan rune = make(chan rune)
	const text string = "Viva La Vida"

	wg.Add(2)

	//emitter thread
	go func() {
		defer func() {
			close(runesChannel)
			wg.Done()
		}()

		for _, char := range text {
			defer func(char rune) {
				runesChannel <- char
			}(char)
		}
	}()

	//receiver thread
	go func() {
		for char := range runesChannel {
			fmt.Print(string(char))
		}
		wg.Done()
	}()

	wg.Wait()
}

