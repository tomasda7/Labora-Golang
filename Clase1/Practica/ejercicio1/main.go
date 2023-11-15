package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese un numero.")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("El numero ingresado es menor a 0.")
	} else if num > 0 {
		fmt.Println("El numero ingresado es mayor a 0.")
	} else {
		fmt.Println("El numero ingresado es igual a 0.")
	}
}
