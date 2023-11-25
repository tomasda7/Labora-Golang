package main

import (
	"fmt"
)

func main() {
	fmt.Println(dnaSequence("ATGCGTAT", "ATATGCGT"))
}

/*
Una cadena de ADN se representa como una secuencia circular de bases (adenina, timina, citosina y guanina) que es única para cada ser vivo, por ejemplo:

Dicha cadena se puede representar como un arreglo de caracteres recorriendola en sentido horario desde la parte superior izquierda:

A T G C G T A T.

Se pide diseñar una clase que represente una secuencia de ADN e incluya un método booleano que nos devuelva **true** si dos cadenas de ADN coinciden.

MUY IMPORTANTE: La secuencia de ADN es cíclica, por lo que puede comenzar en cualquier posición. Por ejemplo, las dos secuencias siguientes coinciden:

**A** T G C G T A T

A T **A** T G C G T
*/

func dnaSequence(sequence1 string, sequence2 string) bool {

	if len(sequence1) > len(sequence2) || len(sequence2) > len(sequence1) {
		return false
	}

	sequence1ToRune := []rune(sequence1)
	sequence2ToRune := []rune(sequence2)


	for i := 0; i < len(sequence1ToRune); i++ {
		match := true
		for j := 0; j < len(sequence1ToRune); j++ {

			if sequence1ToRune[(i+j)%len(sequence1ToRune)] != sequence2ToRune[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}

	return false
}

