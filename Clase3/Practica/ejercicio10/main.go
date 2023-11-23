package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese un numero para averiguar el factorial de su sumatoria:")
	fmt.Scan(&num)

	fmt.Println(factorial(num) * sumatoria(num))
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func sumatoria(num int) int {
	acc := 0

	for i := 1; i <= num; i++ {
		acc += i
	}

	return acc
}
