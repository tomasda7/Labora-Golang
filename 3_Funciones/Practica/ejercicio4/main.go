package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 2; i <= 30; i++ {
		prime_v1(i)
		prime_v2(i)
	}
}

func prime_v1(n int) {
	divCount := 0
	// método de contar cantidad de divisores
	for i := 2; i < n; i++ {
		if (n % i) == 0 {
			divCount++
		}
	}
	isPrime := divCount == 0
	msg := "El número " + strconv.Itoa(n)
	if isPrime {
		msg += " es primo"
	} else {
		msg += " no es primo"
	}
	fmt.Println("v1:", msg)
}

func prime_v2(n int) {
	isPrime := true // método status quo abogadisitico: Asumo que es primo hasta que la realidad me desmuestre lo contrario
	for i := 2; i < n && isPrime; i++ {
		if (n % i) == 0 {
			isPrime = false
		}
	}

	msg := "El número " + strconv.Itoa(n)
	if isPrime {
		msg += " es primo"
	} else {
		msg += " no es primo"
	}
	fmt.Println("v2:", msg)
}

// tienen partes en comun.... }
// ¿hay cosas en común en cada algoritmo? ¿podemos recortar de cada algoritmo la parte que hace propiamente el análisis de saber si es primo y poner la parte que tienen en común en otro lado?
// el prime_v1 tiene relacion con los divisores propios? podriamos escribirlo usando la funcion de divisores propios?
