package main

import (
	"fmt"
	"math/rand"
)

/*
(4) ¿Qué pasaría si quiero una secuencia pero en el orden inverso? 🤔. ¿Acaso es mucho pedir…? Creo que estoy teniendo sed de combinar dos capacidades: por un lado el orden en el que se generan los números: creciente o decreciente y, por otro lado, un predicado que me diga si un número cumple con tal condición.
Se desea implementar una estructura Sequence que tenga como campos un generador y un predicado y que tenga un método llamado Next() que devuelva el próximo número que se genera usando el generador y que cumpla con el predicado.
*/

func main() {
	var intPredicate IntPredicate

	switch rand.Intn(3) {
	case 0:
		intPredicate = Predicate(isInteger)
	case 1:
		intPredicate = Predicate(isPrime)
	case 2:
		intPredicate = Predicate(isEven)
	}

	var intGenerator IntGenerator

	switch rand.Intn(2) {
	case 0:
		intGenerator = AscendingIntGenerator{}
	case 1:
		intGenerator = DescendingIntGenerator{}
	}

	var sequence Sequence = Sequence{current: 0, predicate: intPredicate, generator: intGenerator}

	for i := 1; i <= 30; i++ {
		fmt.Println(sequence.Next())
	}
}

type IntGenerator interface {
	Next(n int) int
}

type AscendingIntGenerator struct {}
type DescendingIntGenerator struct {}

func (g AscendingIntGenerator) Next(n int) int {
	return n + 1
}

func (g DescendingIntGenerator) Next(n int) int {
	return n - 1
}

type IntPredicate interface {
	Fulfill(n int) bool
}

type Predicate func(n int) bool

func (p Predicate) Fulfill(n int) bool {
	return p(n)
}

type Sequence struct {
	current   int
	predicate IntPredicate
	generator IntGenerator
}

func (s *Sequence) Next() int {
	var next int

	for n := s.generator.Next(s.current); n <= 200; n = s.generator.Next(n) {
		if s.predicate.Fulfill(n) {
			next = n
			break
		}
	}
	s.current = next

	return next
}

func isPrime(n int) bool {
	if n <= 0 {
		n = -n
	}

	for i := 2; i <= n/2; i++ {
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
