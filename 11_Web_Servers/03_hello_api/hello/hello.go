package hello

import "strings"

func Greeting(name string, nameList *[]string) string {
	var output string
	for _, n := range *nameList {
		if strings.EqualFold(n, name) {
			output = "¡Nice to see you again " + name + "!"
			return output
		}
	}

	*nameList = append(*nameList, name)
	output = "¡Greetings " + name + "!"

	return output
}

