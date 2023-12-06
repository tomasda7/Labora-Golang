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

	fmt.Println("------------Main begins------------")

	go evensSum(&channel)
	go oddsSum(&channel)

	<-channel
	<-channel
	close(channel)

	fmt.Println("------------Main ends------------")

}


func evensSum(channel *chan bool) {

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
}

func oddsSum(channel *chan bool) {
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
}
