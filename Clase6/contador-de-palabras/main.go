package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "I ate a donut. Then I ate another donut."
	fmt.Println(wordCounter(word))
}

func wordCounter(str string) map[string]int {
	fields := strings.Fields(str)
	mapOfWords := map[string]int{}

	for _, word := range fields {
		count := 0

		for i := 0; i < len(fields); i++ {
			if fields[i] == word {
				count++
			}
		}

		mapOfWords[word] = count
		count = 0
	}

	return mapOfWords
}
