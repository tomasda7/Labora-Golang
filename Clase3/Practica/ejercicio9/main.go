package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Println("Ingrese un numero para saber si es capicúa:")
	fmt.Scan(&num)

	esCapicua(num)
}

func esCapicua(num int) {
	strNum := strconv.Itoa(num)

	runas := []rune(strNum)

	for i, j := 0, len(runas)-1; i < j; i, j = i+1, j-1 {
		runas[i], runas[j] = runas[j], runas[i]
	}

	strInv := string(runas)

	if strNum == strInv {
		fmt.Printf("El numero %d es capicúa.", num)
	} else {
		fmt.Printf("El numero %d no es capicúa.", num)
	}
}
