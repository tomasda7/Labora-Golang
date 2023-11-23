package main

import (
	"fmt"
)

// Comenta que realiza el siguiente programa.

func toInt(cad string) int {
	//se obtiene la longitud de la cadena
	long := len(cad)

	//se declaran 3 variables: 'i' para guardar el indice actual, 'pot' para representar las potencias de 10 y 'res' para almacenar el resultado final.
	i, pot, res := 0, 1, 0

	//se itera mientras la variable 'i' sea menor a la logitud de la cedena.
	for i < long {
		//se extrae el caracter en la posicion 'long-i-1' de la cadena.
		c := cad[long-i-1]
		//se convierte el carcater 'c' en entero, se multiplica por la pontencia de 10 correspondiente y se suma al resultado final.
		res += int(c-'0') * pot
		//se actualiza la variable 'pot' multiplicando su valor por 10.
		pot *= 10
		//el bucle avanza a la siguiente iteracion.
		i++
	}

	//retorna un entero que es la suma total del valor numerico correspondiente a cada caracter de la cadena al terminar de iterarla de derecha a izquierda.
	return res

}

func main() {
	var cadena string
	fmt.Print("Ingrese un número como cadena: ")
	fmt.Scan(&cadena)

	resultado := toInt(cadena)
	fmt.Printf("El resultado como entero es: %d\n", resultado)
}
