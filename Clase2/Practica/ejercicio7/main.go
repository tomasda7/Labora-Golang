package main

import "fmt"

func main() {
	a := 0
	b := 0

	fmt.Println("Ingrese 2 numeros para obtener su maximo comun divisor:")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("El cmd entre %d y %d es %d", a, b, mcd(a,b))
}

//Desarrollar un algoritmo para hallar el máximo común divisor (abreviado como mcd) entre dos números naturales. El máximo común divisor entre dos números es el mayor número que divide a ambos.

func mcd(a int, b int) int {

	//Algoritmo de Euclides
	for b != 0 {
		aux := b
		b = a % b
		a = aux
	}

	return a
}
