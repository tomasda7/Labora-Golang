package main

import "fmt"

func main() {
	var str string
	fmt.Println("Ingrese una palabra para saber si es un palíndromo:")
	fmt.Scan(&str)

	esPalindromo(str)
}

func strInv(str string) string {

	runas := []rune(str)
	var reverseStr []rune

	for i := len(runas) - 1; i >= 0; i-- {
		reverseStr = append(reverseStr, runas[i])
	}

	return string(reverseStr)
}

func esPalindromo(str string) {

	strInvertida := strInv(str)

	if strInvertida == str {
		fmt.Println(str, "es un palíndromo.")
	} else {
		fmt.Println(str, "no es un palíndromo.")
	}

}
