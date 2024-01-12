package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese un numero")
	fmt.Scan(&num)

	if(num % 2 == 0) {
		fmt.Println("El numero ingresado es par.")
	} else  {
		fmt.Println("El numero ingresado es impar.")
	}

}
