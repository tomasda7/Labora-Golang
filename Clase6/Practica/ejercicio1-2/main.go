package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
(1) Se necesita generar dos secuencias de números que sigan la siguiente lógica
Números en orden crecientes (del 1 en adelante)
Números que sean primos (1,2,3,5,7,9,11…)
Se pide implementar cada secuencia en una struct que cumpla con la interfaz.
Por último se pide desarrollar un pequeño programa donde se pueda imprimir los primeros 30 números de una de las dos secuencias según una “tirada de dados”.
Tip: Podemos simular una tirada de dado usando rand.Intn(2) (del paquete math/rand)

(2)Ahora se desea que se imprima un mensaje que diga cuál de las dos secuencias se tomó (según la tirada de dados).
*/


func main() {
	var sequence IntSequence

	if rand.Intn(2) == 0 {
		sequence = &Linear{}
	} else {
		sequence = &Primes{}
	}

	fmt.Println(sequence.Title())
	for i := 0; i <= 30; i++ {
		fmt.Println(sequence.Next())
	}

}

type IntSequence interface {
	Next() int
	Title() string
}

type Linear struct {
	current int
}

func (l *Linear) Title() string {
	return "Linear Sequence"
}

func (l *Linear) Next() int {
	next := l.current
	l.current++

	return next
}


type Primes struct {
	current int
}

func (p *Primes) Title() string {
	return "Primes Sequence"
}

func (p *Primes) Next() int {
	var next int

	for n := p.current + 1 ; n <= 200; n++ {
		if isPrime(n) {
			next = n
			break
		}
	}
	p.current = next

	return next
}

func isPrime(num int) bool {
	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}





