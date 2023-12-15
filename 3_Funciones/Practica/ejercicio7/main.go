package main

import "fmt"

func main() {
	var str string
	fmt.Println("Ingrese una cadena para invertirla:")
	fmt.Scan(&str)

	fmt.Println(strInv(str))

}

// Realice un algoritmo que dado un string lo invierta.
func strInv(str string) string {

	runas := []rune(str)
	var reverseStr []rune

	for i := len(runas) - 1; i >= 0; i-- {
		reverseStr = append(reverseStr, runas[i])
	}

	return string(reverseStr)
}
