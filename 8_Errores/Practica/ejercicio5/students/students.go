package student

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	ID     int                `json:"id"`
	Name   string             `json:"name"`
	Grades map[string]float64 `json:"grades"`
}

func ListStudents(students []Student) {
	if len(students) == 0 {
		fmt.Println("No students registered")
	}

	fmt.Println("---------- Student List ----------")
	for _, s := range students {
		fmt.Printf("id: %d, name: %s, grades: %v \n", s.ID, s.Name, s.Grades)
	}
}

func GenerateID(students []Student) int {
	if len(students) == 0 {
		return 1
	}

	return students[len(students) - 1].ID + 1
}

func AddStudent(students *[]Student, name string, grades map[string]float64) error {
	newStudent := Student{
		Name: name,
		ID: GenerateID(*students),
		Grades: grades,
	}

	for _, s := range *students {
		if strings.EqualFold(s.Name, name) {
			return errors.New("Student entered already exists")
		}
	}

	*students = append(*students, newStudent)

	return nil
}

func SaveStudent(file *os.File, students []Student) {
	bytes, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}

	// position
	_, err = file.Seek(0,0)
	if err != nil {
		panic(err)
	}

	// clean the file
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)

	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	// validates that the file was written
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}


func DeleteGrades(students *[]Student, name string) error {
	for i := range *students {
		if strings.EqualFold((*students)[i].Name,name) {
			(*students)[i].Grades = map[string]float64{}
			return nil
		}
	}

	return errors.New("Student not found")
}

func UpdateGrades(students *[]Student, name string, newGrades map[string]float64) error {

	for i := range *students {
		if strings.EqualFold((*students)[i].Name,name) {
			(*students)[i].Grades = newGrades
			return nil
		}
	}

	return errors.New("Student not found")
}


func CalculateAverageGrade(students []Student, name string) (float64, error) {

	var grades map[string]float64

	for _, s := range students {
		if strings.EqualFold(s.Name, name) {
			grades = s.Grades
			var scores = []float64{}

			for _, score := range grades {
				scores = append(scores, score)
			}

			var totalScores float64
			var length = len(scores)

			for _, score := range scores {
				totalScores += score
			}

			return totalScores / float64(length), nil
		}
	}

	return 0, errors.New("Student not found")
}

