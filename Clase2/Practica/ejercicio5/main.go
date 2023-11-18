package main

import "fmt"

func main() {
fmt.Println("La suma de los numero es", sumSerie())
}

//Redactar un algoritmo que pida al usuario el ingreso de una serie de números reales e imprima por pantalla el resultado de sumarlos. La serie finaliza cuando el usuario ingresa el número cero.

func sumSerie() int {
	serie := []int{}
	n := 0

	fmt.Println("Ingrese una serie de numeros y 0 para finalizar.")
	for {
		fmt.Scan(&n)
		if(n == 0) {
			break
		} else {
			serie = append(serie, n)
		}
	}

	total := 0

	for i := 0; i < len(serie); i++ {
		total += serie[i]
	}

	return total
}
