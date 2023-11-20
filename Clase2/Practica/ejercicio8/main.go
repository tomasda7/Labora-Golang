package main

import "fmt"

func main() {
	a := 0
	b := 0

	fmt.Println("Ingrese 2 numeros para obtener su mínimo común múltiplo:")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Printf("El mcm entre %d y %d es %d", a, b, mcm(a,b))
}

//Desarrollar un algoritmo para hallar el mínimo común múltiplo (abreviado como mcm) entre dos números naturales.
//El mínimo común múltiplo entre dos números es el menor número que es divisible por ambos.
//Ej.: mcm (6,9) = 18, mcm (10,15) = 30, mcm (7,14) = 14, mcm (3,7) = 21

func mcd(a, b int) int {

	//Algoritmo de Euclides
	for b != 0 {
		aux := b
		b = a % b
		a = aux
	}

	return a
}

// Nos servimos de la relacion que existe entre el MCM y el MCD: MCM(a, b) = |a * b| / MCD(a, b).
// Pero para evitar el desbordamiento de enteros en la multiplicación a * b, es mejor dividir primero a por el MCD de a y b, y despues multiplicar el resultado por b.

func mcm(a, b int) int {
	return a / mcd(a, b) * b
}
