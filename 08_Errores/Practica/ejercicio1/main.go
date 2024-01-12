package main

import (
	"errors"
	"fmt"
)

/*
# Ejercicio 1: Validación de entrada

Escribe un programa en Go que solicite al usuario dos números (numerador y denominador), intente realizar la división y maneje el caso en el que el denominador sea cero. Si el denominador es cero, imprime un mensaje de error indicando que la división no es posible. En caso contrario, imprime el resultado de la división con dos decimales.
*/

func main() {
	var num, deno int

	fmt.Println("Enter the numerator and denominator to divide:")
	fmt.Scan(&num)
	fmt.Scan(&deno)

	if result, err := divideBy(num,deno); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f", result)
	}
}

func divideBy(n, d int) (float64, error) {

	if d == 0 {
		return 0, errors.New("is not possible to divide by zero.")
	}

	result := float64(n) / float64(d)

	return result, nil
}
