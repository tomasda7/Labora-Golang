package main

import (
	"testing"
)

func TestAddNoteWorksOk(t *testing.T) {
	type Test struct {
		notesToAdd    []int
		studentName   string
		expectedNotes []int
	}

	var tests []Test = []Test{
		{
			notesToAdd:    []int{6, 10},
			studentName:   "pepe",
			expectedNotes: []int{6, 10},
		},
		{
			notesToAdd:    []int{6, 10, 6, 8},
			studentName:   "lopez",
			expectedNotes: []int{6, 10, 6},
		},
		{
			notesToAdd:    []int{6},
			studentName:   "pepe",
			expectedNotes: []int{6, 10, 6},
		},
		{
			notesToAdd:    []int{5, 11, 8},
			studentName:   "miguel",
			expectedNotes: []int{5, 8},
		},
	}

	var notesByName map[string][]int = make(map[string][]int)

	for _, test := range tests {
		for _, note := range test.notesToAdd {
			AddNote(notesByName, test.studentName, note)
		}

		computedNotes, exists := notesByName[test.studentName]
		if !exists {
			t.Error("Se esparaba encontrar una nota")
		}
		if !areSlicesEquals(test.expectedNotes, computedNotes) {
			t.Errorf("Notas esperadas(%d) difiere de notas computadas(%d)", test.expectedNotes, computedNotes)
		}
	}
}

func TestAverageWorksOk(t *testing.T) {
	type Test struct {
		notes           []int
		studentName     string
		expectedAverage float64
	}

	var tests []Test = []Test{
		{
			notes:           []int{6, 8},
			studentName:     "pepe",
			expectedAverage: 7,
		},
	}

	var notesByName map[string][]int = make(map[string][]int)
	for _, test := range tests {
		for _, note := range test.notes {
			AddNote(notesByName, test.studentName, note)
		}

		computedAverage := calculateAverages(notesByName)
		if computedAverage[test.studentName] != test.expectedAverage {
			t.Errorf("Promedio esperado(%f) difiere de promedio computado(%f)", test.expectedAverage, computedAverage[test.studentName])
		}
	}
}

func areSlicesEquals(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {

		if left[i] != right[i] {
			return false
		}
	}

	return true
}
