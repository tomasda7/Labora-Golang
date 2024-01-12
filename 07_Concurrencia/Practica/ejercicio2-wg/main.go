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
	var wg sync.WaitGroup

	var evensTime time.Duration
	var oddsTime time.Duration

	wg.Add(2)

	go func ()  {
		evensTime = evensSum(&wg)
	}()
	go func ()  {
		oddsTime = oddsSum(&wg)
	}()

	wg.Wait()
	fmt.Println("-------- Final Results --------")
	if oddsTime < evensTime {
		fmt.Println("oddsSum Wins!")
	} else if evensTime < oddsTime {
		fmt.Println("evensSum Wins!")
	} else {
		fmt.Println("It is a tie!")
	}
}


func evensSum(wg *sync.WaitGroup) time.Duration {

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
	return end
}

func oddsSum(wg *sync.WaitGroup) time.Duration {

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
	return end
}
