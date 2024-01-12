package main

import (
	"testing"
)

 func TestLastIndex(t *testing.T) {
	type Test struct {
		n           int // cantidad de primeros impares a sumar
		expectedSum int // resultado esperado
	}

	tests := []Test{
		{
			n:           0,
			expectedSum: 0,
		},
		{
			n:           1,
			expectedSum: 1,
		},
		{
			n:           2,
			expectedSum: 4,
		},
		{
			n:           5,
			expectedSum: 25,
		},
	}

	for _, test := range tests {
		computedSum := sumOddsThree(test.n)
		if computedSum != test.expectedSum {
			t.Errorf("sumOdds(%d) = %d, expected was %d", test.n, computedSum, test.expectedSum)
		}
	}
}
// go test -run TestLastIndex

func BenchmarkSumOddsOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumOddsOne(b.N)
	}
}

func BenchmarkSumOddsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumOddsTwo(b.N)
	}
}

func BenchmarkSumOddsThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumOddsThree(b.N)
	}
}
// go test -bench .
