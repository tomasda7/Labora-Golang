package main

import "fmt"

func main() {
	var str string
	println("Ingrese una cadena de texto para transformarla:")
	fmt.Scan(&str)

	fmt.Println(replaceVocal(str))
}

func replaceVocal(str string) string {
	var result string

	for _, value := range str {
		switch value {
		case 'a':
			result += "4"
		case 'e':
			result += "3"
		case 'i':
			result += "1"
		case 'o':
			result += "0"
		case 'u':
			result += "6"
		default:
			result += string(value)
		}
	}

	return result
}
