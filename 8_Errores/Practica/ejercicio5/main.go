package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	student "github.com/tomasda7/go-grades-crud/students"
)

func main() {

	file, err := os.OpenFile("students.json", os.O_RDWR | os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var students []student.Student

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &students)
		if err != nil {
			panic(err)
		}
	} else {
		students = []student.Student{}
	}


	// UI
	if(len(os.Args) < 2) {
		printUsage()
	} else {
		switch os.Args[1] {
		case "list":
			student.ListStudents(students)
		case "add":
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Enter the student name:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("How many subjects do you want to add?")
			numOfSubjectsStr, _ := reader.ReadString('\n')
			numOfSubjectsStr = strings.TrimSpace(numOfSubjectsStr)

			numOfSubjects, err := strconv.Atoi(numOfSubjectsStr)
			if err != nil {
			fmt.Println("Invalid number. Please enter a valid number.")
			return
			}

			grades := make(map[string]float64)

			for i := 0; i < numOfSubjects; i++ {
				fmt.Printf("Enter the name of subject %d:\n", i+1)
				subject, _ := reader.ReadString('\n')
				subject = strings.TrimSpace(subject)

				fmt.Printf("Enter the grade for %s:\n", subject)
				scoreStr, _ := reader.ReadString('\n')
				scoreStr = strings.TrimSpace(scoreStr)

				score, err := strconv.ParseFloat(scoreStr, 64)

				if err != nil {
					fmt.Printf("Invalid grade for %s. Please enter a valid number.\n", subject)
					i--
					continue
				}

				grades[subject] = score
			}

			err = student.AddStudent(&students, name, grades)
			if err != nil {
				fmt.Println("Error trying to add a new student:", err)
			}

			student.SaveStudent(file, students)
		case "deleteGrades":
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Enter the student name:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			err := student.DeleteGrades(&students, name)
			if err != nil {
				fmt.Println("Error trying to delete grades:", err)
			}

			student.SaveStudent(file, students)
		case "updateGrades":
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Enter a student name to update their grades:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("How many subjects do you want to add?")
			numOfSubjectsStr, _ := reader.ReadString('\n')
			numOfSubjectsStr = strings.TrimSpace(numOfSubjectsStr)

			numOfSubjects, err := strconv.Atoi(numOfSubjectsStr)
			if err != nil {
			fmt.Println("Invalid number. Please enter a valid number.")
			return
			}

			grades := make(map[string]float64)

			for i := 0; i < numOfSubjects; i++ {
				fmt.Printf("Enter the name of subject %d:\n", i+1)
				subject, _ := reader.ReadString('\n')
				subject = strings.TrimSpace(subject)

				fmt.Printf("Enter the grade for %s:\n", subject)
				scoreStr, _ := reader.ReadString('\n')
				scoreStr = strings.TrimSpace(scoreStr)

				score, err := strconv.ParseFloat(scoreStr, 64)

				if err != nil {
					fmt.Printf("Invalid grade for %s. Please enter a valid number.\n", subject)
					i--
					continue
				}

				grades[subject] = score
			}

			err = student.UpdateGrades(&students, name, grades)
			if err != nil {
				fmt.Println("Error trying to update grades:", err)
			}

			student.SaveStudent(file, students)
		case "averageGrade":
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Enter a student's name to calculate their average grade:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			averageGrade, err := student.CalculateAverageGrade(students, name)
			if err != nil {
				fmt.Println("Error trying to calculate average grade:", err)
			} else {
				fmt.Printf("The average grade of %s is %.1f", name, averageGrade)
			}

		default:
			printUsage()
		}
	}
}

func printUsage() {
	fmt.Println("Usage: go run main.go [list|add|updateGrades|deleteGrades|averageGrade]")
}



