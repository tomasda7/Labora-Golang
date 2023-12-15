package main

import (
	"fmt"
	"time"
)

/*
# Ejercicio 2

Escriba un programa en donde haya dos gorutinas donde una suma los primeros 100 números pares y la otra los 100 primeros números impares. Se desea saber cual gorutina “gana”, es decir, quien termina termina primero.

Para sincronizar se puede usar grupos de espera o canales. Preferentemente podés realizar el ejercicio con ambos mecanismos para que te vayas familiarizando.
*/


func main() {
	channel := make(chan bool)

	var oddsTime time.Duration
	var evensTime time.Duration


	go func ()  {
		oddsTime = oddsSum(&channel)
	}()

	go func ()  {
		evensTime = evensSum(&channel)
	}()

	<-channel
	<-channel
	close(channel)

	fmt.Println("-------- Final Results --------")
	if oddsTime < evensTime {
		fmt.Println("oddsSum Wins!")
	} else if evensTime < oddsTime {
		fmt.Println("evensSum Wins!")
	} else {
		fmt.Println("It is a tie!")
	}
}


func evensSum(channel *chan bool) time.Duration {
	begin := time.Now()

	var total int
	for i := 0; i <= 10000000; i++ {
		if i % 2 == 0 {
			total += i
		}
	}

	end := time.Since(begin)
	fmt.Println("evens sum ended in", end)
	*channel <- true
	return end
}

func oddsSum(channel *chan bool) time.Duration {
	begin := time.Now()

	var total int
	for i := 0; i <= 10000000; i++ {
		if i % 2 != 0 {
			total += i
		}
	}

	end := time.Since(begin)
	fmt.Println("odds sum ended in", end)
	*channel <- true
	return end
}
