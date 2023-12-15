package main

import "fmt"

func main() {
	num := 0

	fmt.Println("Ingrese un numero del 1 al 100:")
	fmt.Scan(&num)
	fizzBuzz(num)
}

//Ejercicio FizzBuzz en su forma clásica es el siguiente: "Para cada número del 1 al 100, imprime 'Fizz' si el número es divisible por 3, 'Buzz' si es divisible por 5 y 'FizzBuzz' si es divisible por ambos. Si no es divisible ni por 3 ni por 5, simplemente imprime el número."

func fizzBuzz(num int) {
	if num % 3 == 0 && num % 5 == 0 {
		fmt.Println("FizzBuzz")
	} else if num % 3 == 0 {
		fmt.Println("Fizz")
	} else if num % 5 == 0  {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}
}
