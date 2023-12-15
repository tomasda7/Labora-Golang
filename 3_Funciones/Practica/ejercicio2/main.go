// Ejercicio 2:
package main

import "fmt"

func main() {
    var a int
    var b int

    fmt.Println("Ingrese dos numeros para saber si son 'numeros amigos': ")
    fmt.Scan(&a)
    fmt.Scan(&b)

    if sonAmigos(a, b) {
		fmt.Printf("%d y %d son numeros amigos", a, b)
	} else {
		fmt.Printf("%d y %d no son numeros amigos", a, b)
	}
}

/*
2.Proponga una solución que use funciones bien diseñadas (solución modular) al problema de determinar si dos números naturales son amigos. O sea la ideas es pensar en cómo puede reutilizar la función del ejercicio anterior para si un número es perfecto.
Un número es amigo de otro cuando la suma de sus divisores propios es igual al otro número.
*/

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

func sonAmigos(numeroA int, numeroB int) bool {
    sumA := sumaDivisores(numeroA)
    sumB := sumaDivisores(numeroB)

    return sumA == numeroB && sumB == numeroA
}


