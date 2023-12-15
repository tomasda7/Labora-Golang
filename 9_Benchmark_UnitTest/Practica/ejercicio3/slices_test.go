package slices

import (
	"fmt"
	"testing"
)

func TestRotareRightWorks(t *testing.T) {
	str := "ABC"
	expected := "CAB"
	generated := RotateRight(str)
	if generated != expected {
		t.Errorf("No son iguales, generated: %s, expected: %s", generated, expected)
	}

	str = "ABCD"
	expected = "DABC"
	generated = RotateRight(str)
	if generated != expected {
		t.Errorf("No son iguales, generated: %s, expected: %s", generated, expected)
	}

	str = "MNO"
	expected = "OMN"
	generated = RotateRight(str)
	fmt.Println(str, generated)
	if generated != expected {
		t.Errorf("No son iguales, generated: %s, expected: %s", generated, expected)
	}
}

func TestRotareRightWorksTDT(t *testing.T) {
	type Test struct {
		input    string
		expected string
	}

	var tests []Test = []Test{
		{
			input:    "",
			expected: "",
		},
		{
			input:    "AA",
			expected: "AA",
		},
		{
			input:    "ABC",
			expected: "CAB",
		},
		{
			input:    "MNO",
			expected: "OMN",
		},
		{
			input:    "xyza",
			expected: "axyz",
		},
	}

	for _, test := range tests {
		generated := RotateRightVersion4(test.input)
		if generated != test.expected {
			t.Errorf("No son iguales, generated: %s, expected: %s", generated, test.expected)
		}
	}
}

func TestRotareRightTimesWorks(t *testing.T) {
	type Test struct {
		input    string
		times    int
		expected string
	}

	var tests []Test = []Test{
		{
			input:    "",
			expected: "",
			times:    1,
		},
		{
			input:    "AA",
			expected: "AA",
			times:    1,
		},
		{
			input:    "ABC",
			expected: "CAB",
			times:    1,
		},
		{
			input:    "ABC",
			expected: "BCA",
			times:    2,
		},
		{
			input:    "ABCD",
			expected: "CDAB",
			times:    2,
		},
		{
			input:    "xyza",
			expected: "zaxy",
			times:    2,
		},
	}

	for _, test := range tests {
		generated := RotateRightTimes(test.input, test.times)
		if generated != test.expected {
			t.Errorf("No son iguales, generated: %s, expected: %s", generated, test.expected)
		}
	}
}

func TestRotaresWorks(t *testing.T) {
	type Test struct {
		input                string
		expectedRotatedRight string
		expectedRotatedLeft  string
	}

	var tests []Test = []Test{
		{
			input:                "",
			expectedRotatedRight: "",
			expectedRotatedLeft:  "",
		},
		{
			input:                "AA",
			expectedRotatedRight: "AA",
			expectedRotatedLeft:  "AA",
		},
		{
			input:                "ABC",
			expectedRotatedRight: "CAB",
			expectedRotatedLeft:  "BCA",
		},
		{
			input:                "MNO",
			expectedRotatedRight: "OMN",
			expectedRotatedLeft:  "NOM",
		},
		{
			input:                "xyza",
			expectedRotatedRight: "axyz",
			expectedRotatedLeft:  "yzax",
		},
	}

	for _, test := range tests {
		generatedRotatedRight := RotateChainRight(test.input)
		if generatedRotatedRight != test.expectedRotatedRight {
			t.Errorf("RotateChainRight: No son iguales, generated: %s, expected: %s", generatedRotatedRight, test.expectedRotatedRight)
		}
		generatedRotatedLeft := RotateChainLeft(test.input)
		if generatedRotatedLeft != test.expectedRotatedLeft {
			t.Errorf("RotateChainLeft: No son iguales, generated: %s, expected: %s", generatedRotatedRight, test.expectedRotatedRight)
		}
	}
}

func RotateChainLeft(chain string) string {
	var result string
	n := len(chain)

	for i := 0; i < n; i++ {
		result += string(chain[(n+i+1)%n])
	}

	return result
}

func RotateChainRight(chain string) string {
	var result string
	n := len(chain)

	for i := 0; i < n; i++ {
		result += string(chain[(n+i-1)%n])
	}

	return result
}

// go test -v  -run .
