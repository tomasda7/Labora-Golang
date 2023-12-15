package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
# Ejercicio 3: Conversión de tipos con manejo de errores

Escribe una función que convierta un string a un número entero (`int`). La función debe devolver un error si el string no puede convertirse a un número. Prueba tu función con diferentes strings, incluyendo algunos que no sean números. Usa `panic` y `recover` para manejar los errores
*/

func main() {
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	invalidInput := "F"
	//validInput := "1234"

	if num, err := stringToInt(invalidInput); err != nil {
		panic(err)
	} else {
		fmt.Println("Result:", num)
	}
}


func stringToInt(str string) (int,error) {
	if number, err := strconv.Atoi(str); err != nil {
		return 0, errors.New(err.Error())
	} else {
		return number, nil
	}
}
