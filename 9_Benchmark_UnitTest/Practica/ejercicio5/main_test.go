package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type Test struct {
	s []int
}

var test = []Test{
	{
		s: rand.Perm(10),
	},
	{
		s: rand.Perm(100),
	},
	{
		s: rand.Perm(1000),
	},
	{
		s: rand.Perm(10000),
	},
	{
		s: rand.Perm(100000),
	},
	{
		s: rand.Perm(1000000),
	},
	{
		s: rand.Perm(10000000),
	},
	{
		s: rand.Perm(100000000),
	},
}

func BenchmarkSumaSecuencial(b *testing.B) {
	for _, value := range test {
		b.Run(fmt.Sprintf("SECUENCIAL:"), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sumaSecuencial(value.s)
			}
		})
	}
}
func BenchmarkSumaConcurrente(b *testing.B) {
	for _, value := range test {
		b.Run(fmt.Sprintf("CONCURRENTE:"), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sumaConcurrente(value.s)
			}
		})
	}
}
