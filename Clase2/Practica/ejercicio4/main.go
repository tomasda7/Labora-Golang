package main

import "fmt"

func main() {
fmt.Println("El resultado de la suma es", sumN())
}

//Realice un algoritmo para sumar los primeros "n" números naturales, siendo "n" un valor que ingresa el usuario por teclado durante la ejecución del algoritmo

func sumN() int {
	var n int
	fmt.Println("Ingrese un numero:")
	fmt.Scan(&n)

	total := 0

	for i := 1; i <= n; i++ {
		total += i
	}

	return total
}
