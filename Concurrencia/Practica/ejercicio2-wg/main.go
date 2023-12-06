package main

import (
	"fmt"
	"sync"
	"time"
)

/*
# Ejercicio 2

Escriba un programa en donde haya dos gorutinas donde una suma los primeros 100 números pares y la otra los 100 primeros números impares. Se desea saber cual gorutina “gana”, es decir, quien termina termina primero.

Para sincronizar se puede usar grupos de espera o canales. Preferentemente podés realizar el ejercicio con ambos mecanismos para que te vayas familiarizando.
*/

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("------------Main begins------------")

	wg.Add(2)
	go evensSum(wg)
	go oddsSum(wg)

	wg.Wait()
	fmt.Println("------------Main ends------------")

}


func evensSum(wg *sync.WaitGroup) {

	begins := time.Now()

	var total int
	for i := 0; i <= 10000000; i++ {
		if i % 2 == 0 {
			total += i
		}
	}

	end := time.Since(begins)

	fmt.Println("evens sum ended in", end)
	wg.Done()
}

func oddsSum(wg *sync.WaitGroup) {

	begins := time.Now()

	var total int
	for i := 0; i <= 10000000; i++ {
		if i % 2 != 0 {
			total += i
		}
	}

	end := time.Since(begins)

	fmt.Println("odds sum ended in", end)
	wg.Done()
}
