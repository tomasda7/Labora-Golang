package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
3) De forma más general se desea tener una secuencia que me dé números que cumplan una condición (predicado).
Dicha condición será implementada en estructuras (o cualquier otro tipo de dato)  que tengan el un método con la siguiente firma Fulfill(n int) bool. Fulfill es un término en inglés que comúnmente se emplea para afirmar que algo cumple una condición.
Además de tener tipos de datos que tengan una función para implementar condiciones de:
Números cualquier sea (para secuencias de enteros lineales)
Números primos
Ahora quiero que agreguen un nuevo tipo que tenga una función que implemente la condición de
Números pares (2,4,6…)
*/

func main() {

	var sequence IntSequence

	switch rand.Intn(3) {
	case 0:
		sequence = &Sequence{current: 0, predicate: Predicate(isInteger), title: "Linear Sequence"}
	case 1:
		sequence = &Sequence{current: 1, predicate: Predicate(isPrime), title: "Primes Sequence"}
	case 2:
		sequence = &Sequence{current: 0, predicate: Predicate(isEven), title: "Evens Sequence"}
	}

	fmt.Println(sequence.Title())
	for i := 0; i < 30; i++ {
		fmt.Println(sequence.Next())
	}

}

type IntPredicate interface {
	Fulfill(n int) bool
}

type Predicate func(n int) bool

func (p Predicate) Fulfill(n int) bool {
	return p(n)
}

type IntSequence interface {
	Next() int
	Title() string
}

type Sequence struct {
	current int
	predicate IntPredicate
	title string
}

func (s *Sequence) Title() string {
	return s.title
}

func (s *Sequence) Next() int {
	var next int

	for n := s.current + 1 ; n <= 200; n++ {
		if s.predicate.Fulfill(n) {
			next = n
			break
		}
	}
	s.current = next

	return next
}

func isPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isInteger(n int) bool {
	return true
}

func isEven(n int) bool {
	return n%2 == 0
}
