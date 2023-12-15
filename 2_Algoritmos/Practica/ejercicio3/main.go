package main

import "fmt"

func main() {
	nums3 := []int{7, 10, 15}
	nums10 := []int{15,20,30,3,65,6,89,32,5,77}

	fmt.Println(sumarNaturales(nums3))
	fmt.Println(sumarNaturales(nums10))
}

//Realice un algoritmo para sumar los primeros 3 números naturales. Y luego un algoritmo para sumar los primeros 10 números naturales

func sumarNaturales(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total
}
