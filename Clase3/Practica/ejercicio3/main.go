package main

import (
	"fmt"
	"math"
)

func main() {
  	var numero int

  	fmt.Println("Ingrese un numero para saber si es primo:")
  	fmt.Scan(&numero)

	if esPrimo(numero) {
	  fmt.Printf("%d es primo", numero)
	} else {
	  fmt.Printf("%d no es primo", numero)
  	}
}

//Escriba un algoritmo para determinar si un número es primo. Recordar que número primo es aquel que es divisible por solo por 1 y por si mismo. ¿ Se puede utilizar alguna función desarrollada en los ejercicios anteriores?

//Ejercicio 3
func esPrimo(num int) bool {

	//recorre hasta la raiz cuadrada del numero
	for i := 2; i < int(math.Sqrt(float64(num))) + 1; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}
