package main

import (
	"fmt"
)

/*
# Ejercicio 2

Escriba un programa en donde haya dos gorutinas donde una suma los primeros 100 números pares y la otra los 100 primeros números impares. Se desea saber cual gorutina “gana”, es decir, quien termina termina primero.

Para sincronizar se puede usar grupos de espera o canales. Preferentemente podés realizar el ejercicio con ambos mecanismos para que te vayas familiarizando.
*/

var channel chan bool = make(chan bool)

func main() {

	fmt.Println("------------Main begins------------")


	go oddsSum()
	go evensSum()

	<-channel
	fmt.Println("------------Main ends------------")

}


func evensSum() {
	var total int
	for i := 0; i <= 1000; i++ {
		if i % 2 == 0 {
			total += i
		}
	}

	fmt.Println("evens sum wins")
	channel <- true
}

func oddsSum() {

	var total int
	for i := 0; i <= 1000; i++ {
		if i % 2 != 0 {
			total += i
		}
	}

	fmt.Println("odds sum wins")
	channel <- true
}
