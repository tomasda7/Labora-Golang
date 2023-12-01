package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var nums Numbers

	nums.Add(10)
	nums.Add(20)
	nums.Add(30)
	nums.Add(40)
	nums.Add(80)
	nums.Add(70)

	fmt.Println(nums.elements, nums.Length())

	toRemoveIndex := 5
	if nums.isInBoundaries(toRemoveIndex) {
		fmt.Println("Remove one:", nums.RemoveAt(toRemoveIndex))
	} else {
		fmt.Printf("Index %d out of boundaries.\n", toRemoveIndex)
	}

	fmt.Println(nums.elements, nums.Length())

	toSetIndex := 4
	if nums.isInBoundaries(toSetIndex) {
		fmt.Println("Set one:", nums.Set(50,toSetIndex))
	} else {
		fmt.Printf("Index %d out of boundaries.\n", toSetIndex)
	}

	fmt.Println(nums.elements, nums.Length())

	fmt.Println("ToString:", nums.ToString())
}

type Numbers struct {
	elements []int
}

func (n *Numbers) Add(value int) int {
	n.elements = append(n.elements, value)
	return value
}

func (n *Numbers) isInBoundaries(index int) bool {
	if n.Length() > index {
		return true
	} else {
		return false
	}
}

func (n *Numbers) RemoveAt(index int) bool {

	var initalLength int = len(n.elements)

	copy(n.elements[index:], n.elements[index+1:])
	// fmt.Println("RemoveAt moves n.elements[index+1:] left:", n.elements)

	n.elements[len(n.elements)-1] = 0
	// fmt.Println("RemoveAt 'deletes' the element by index:", n.elements)

	n.elements = n.elements[:len(n.elements)-1]
	// fmt.Println("RemoveAt truncates the slice:", n.elements)

	if initalLength > len(n.elements) {
		return true
	} else {
		return false
	}
}

func (n *Numbers) Length() int {
	return len(n.elements)
}

func (n *Numbers) Set(value int, index int) bool {

	initialValue := n.elements[index]

	n.elements[index] = value

	if n.elements[index] != initialValue {
		return true
	} else {
		return false
	}
}

func (n Numbers) ToString() string {
	numbers := n.elements

	var strNums []string //slice

	for _, num := range numbers {
		strNums = append(strNums, strconv.Itoa(num))
	}

	str := strings.Join(strNums, ", ")

	return str
}
