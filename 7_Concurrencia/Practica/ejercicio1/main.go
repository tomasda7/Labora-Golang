package main

import (
	"fmt"
	"time"
)

/*
# Ejercicio 1
Usando la técnica del sleep sort, hagamos un programa que imprima los primeros 10 números enteros en orden ascendiente.
¿Es posible hacerlo en orden descendiente?
*/

func main() {
	fmt.Println("------------Main begins------------")

	go descendingCount()
	go ascendingCount()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("------------Main ends------------")
}

func ascendingCount() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func descendingCount() {
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}
}
