package main

import "fmt"

func main() {
	var input int
	fmt.Println("Ingrese un numero para obtener la suma de sus divisores propios:")
	fmt.Scan(&input)
	ownDivisors := getDivisors(input)
	fmt.Println(ownDivisors)
	fmt.Println(sumOfDivisors(ownDivisors))
}

func getDivisors(num int) []int {
	var divisores []int

	for i := 1; i < num; i++ {
		if num%i == 0 && i != num {
			divisores = append(divisores, i)
		}
	}

	return divisores
}

func sumOfDivisors(array []int) int {
	sumaDivisoresPosibles := 0

	for _, value := range array {
		sumaDivisoresPosibles += value
	}

	return sumaDivisoresPosibles
}
