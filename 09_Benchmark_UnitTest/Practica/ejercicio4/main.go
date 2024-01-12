package main

/*
Se pide implementar al menos dos formas de sumar los primeros “n” números impares (una de ellas deberías tener por el ejercicio 1 de esta guía). Y luego comparar cual es más eficiente.
*/

func sumOddsOne(n int) int {
	var acc int
	odd := 1

	for i := 0; i < n; i++ {
		acc += odd
		odd += 2
	}

	return acc
}

func sumOddsTwo(n int) int {
	sum := 0

	for i := 1; i <= n; i++ {
		sum += 2*i - 1
	}

	return sum
}

func sumOddsThree(n int) int {
	return n * n
}
