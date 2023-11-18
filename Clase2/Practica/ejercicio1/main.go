package main

import "fmt"

func main() {
	nums3 := []int{10,3,15}
	nums10 := []int{15,20,30,3,65,6,89,32,5,77}
	fmt.Println(encontrarMayor(nums3))
	fmt.Println(encontrarMayor(nums10))
}

//Realizar un algoritmo para determinar el mayor de 3 números. Y luego para determinar el mayor de 10 numeros.

func encontrarMayor(nums []int) int {
	mayor := nums[0]
	for i := 0; i < len(nums); i++ {
		if mayor < nums[i] {
			mayor = nums[i]
		}
	}

	return mayor
}

