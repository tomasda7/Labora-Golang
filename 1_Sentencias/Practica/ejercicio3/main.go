package main

import "fmt"

func main() {
	var twodays int
	var yesterday int
	var today int

	fmt.Println("Ingrese la temperatura de hace dos dias")
	fmt.Scan(&twodays)
	fmt.Println("Ingrese la temperatura de ayer")
	fmt.Scan(&yesterday)
	fmt.Println("Ingrese la temperatura de hoy")
	fmt.Scan(&today)

	var total int = twodays + yesterday + today
	var average int = total / 3

	fmt.Printf("El total de las temperaturas es %d° \n", total)
	fmt.Printf("El promedio de las temperaturas es %d°", average)
}
