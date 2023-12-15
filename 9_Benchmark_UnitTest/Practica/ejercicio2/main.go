package main

import (
	"errors"
	"fmt"
)

func main() {
	var notesByName map[string][]int = make(map[string][]int)

	exit := false
	for !exit {
		displayMenu()
		operation, err := enterOperation()
		if err != nil {
			fmt.Println("Error introduciendo operación, el error fue:", err)
			continue
		}
		switch operation {
		case 1:
			name, note, err := enterNameAndNote()
			if err != nil {
				fmt.Println("Error introduciendo nombre y/o nota, el error fue:", err)
				continue
			}
			AddNote(notesByName, name, note)
		case 2:
			averageByName := calculateAverages(notesByName)
			fmt.Println(averageByName)
		case 3:
			exit = true
		}
		fmt.Println("algo")
	}
}

func AddNote(notesByName map[string][]int, name string, note int) error {
	// agregar código para verificar que no se estén violando los requisitos de agregar nota!
	// nota debe ser un valor entre 1 y 10
	// no debe haber más de tres notas por estudiante!
	if note < 1 || note > 10 {
		return errors.New("la nota debe ser un valor entre 1 y 10")
	} else if len(notesByName[name]) < 3 {
		notesByName[name] = append(notesByName[name], note)
		return nil
	} else {
		return errors.New("no se pueden agregar más de tres notas por estudiante")
	}
}

func displayMenu() {
	fmt.Println("Sistema alumnex 2.0")
	fmt.Println("====================")

	fmt.Printf("[%-30s]\n", "1. Ingreso de nota")
	fmt.Printf("[%-30s]\n", "2. Ver promedios")
	fmt.Printf("[%-30s]\n", "3. Salir")
}

func enterOperation() (int, error) {
	var operation int
	_, err := fmt.Scanf("%d", &operation)
	if err != nil {
		return 0, err
	}
	return operation, nil
}

func enterNameAndNote() (string, int, error) {
	fmt.Print("Nombre:")
	var name string
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		return "", 0, err
	}

	fmt.Print("Nota:")
	var note int
	_, err = fmt.Scanf("%d", &note)
	if err != nil {
		return "", 0, err
	}
	return name, note, nil
}

func calculateAverages(notesByName map[string][]int) map[string]float64 {
	var averageByName map[string]float64 = make(map[string]float64)
	for name, notes := range notesByName {
		notesSum := 0
		notesCount := 0
		for _, note := range notes {
			notesSum += note
			notesCount++
		}
		average := float64(notesSum) / float64(notesCount)
		averageByName[name] = average
	}
	return averageByName
}
