package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese un numero para saber si es perfecto:")
	fmt.Scan(&num)

	if esPerfecto(num) {
		fmt.Printf("%d es un numero perfecto", num)
	} else {
		fmt.Printf("%d no es un numero perfecto", num)
	}
}

//Desarrolle una función para determinar si un número natural (n > 0) es perfecto.
//Un número es perfecto si la suma de los divisores propios da el el mismo numero.
//Los divisores propios de un número son todos sus divisores sin contar el mismo número.

func sumaDivisores(numero int) int {
	sumaDivisoresPosibles := 0

	for i := 1; i <= numero/2; i++ {
		if numero%i == 0 {
			// si correctamente es un divisor, lo sumamos a la acumulación
			sumaDivisoresPosibles += i
		}
	}

	return sumaDivisoresPosibles
}

func esPerfecto(numero int) bool {
	//Acumulador de los divisores
	var sumaDeDivisoresPosibles int
	//Iteramos hasta la mitad(divisores posibles)
	//comparamos la acumulación y el numero original
	sumaDeDivisoresPosibles = sumaDivisores(numero)

	return sumaDeDivisoresPosibles == numero
}
