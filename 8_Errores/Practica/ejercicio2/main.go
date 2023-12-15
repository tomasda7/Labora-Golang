package main

import (
	"errors"
	"fmt"
)

/*
# Ejercicio 2: Validación de entrada

Crea una función que acepte un string y devuelva un error si el string está vacío o si es demasiado largo (por ejemplo, más de 100 caracteres).
*/

func main() {
	inputVoid := ""
	//inputLong := "asdjhaskjdhjksahdjskahdjsahdjkhdsjkdsadssddddsd"

	if str, err := strValidator(inputVoid); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(str)
	}
}


func strValidator(str string) (string,error) {
	if str == "" {
		return "", errors.New("string too short.")
	} else if len(str) > 40 {
		return "", errors.New("the limit is 40 characters.")
	}

	return str, nil
}
