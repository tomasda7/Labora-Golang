package main

import "fmt"

func main() {
	var longitud int
	var anchura int

	fmt.Println("Ingrese la longitud")
	fmt.Scan(&longitud)
	fmt.Println("Ingrese la anchura")
	fmt.Scan(&anchura)

	var area int = longitud * anchura
	var perimetro int = 2 * (longitud + anchura)

	fmt.Printf("El area del rectangulo es %d \n", area)
	fmt.Printf("El perimetro del rectangulo es %d", perimetro)
}

/*
Descripción del Problema:
Escribe un programa que calcule el área y el perímetro de un rectángulo. El programa debe pedir al usuario que introduzca la longitud y la anchura del rectángulo y luego calcular y mostrar el área y el perímetro.
Especificaciones:
Solicita al usuario que ingrese la longitud y la anchura del rectángulo.
Calcula el área del rectángulo (longitud * anchura).
Calcula el perímetro del rectángulo (2 * (longitud + anchura)).
Muestra los resultados (área y perímetro) al usuario.
*/
