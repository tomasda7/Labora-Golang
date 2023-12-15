package odds

import (
	"testing"
)

 func TestSumFirstOdds(t *testing.T) {
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
		computedSum := sumFirstOdds(test.n)
		if computedSum != test.expectedSum {
			t.Errorf("sumOdds(%d) = %d, expected was %d", test.n, computedSum, test.expectedSum)
		}
	}
}
// go test -run TestSumFirstOdds
