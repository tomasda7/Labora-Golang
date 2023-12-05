package main

import (
	"fmt"
	"sync"
)

/*
# Ejercicio 2

Escriba un programa en donde haya dos gorutinas donde una suma los primeros 100 números pares y la otra los 100 primeros números impares. Se desea saber cual gorutina “gana”, es decir, quien termina termina primero.

Para sincronizar se puede usar grupos de espera o canales. Preferentemente podés realizar el ejercicio con ambos mecanismos para que te vayas familiarizando.
*/

func main() {
	wg := &sync.WaitGroup{}

	fmt.Println("------------Main begins------------")

	wg.Add(1)
	go evensSum(wg)
	go oddsSum(wg)

	wg.Wait()
	fmt.Println("------------Main ends------------")

}


func evensSum(wg *sync.WaitGroup) {
	var total int
	for i := 0; i <= 1000; i++ {
		if i % 2 == 0 {
			total += i
		}
	}

	fmt.Println("evens sum wins")
	wg.Done()
}

func oddsSum(wg *sync.WaitGroup) {

	var total int
	for i := 0; i <= 1000; i++ {
		if i % 2 != 0 {
			total += i
		}
	}

	fmt.Println("odds sum wins")
	wg.Done()
}
