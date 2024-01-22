package main

import (
	"fmt"
	"strings"
)

// reverse words in a line
func reverseWords(line string) string {
	// splits line at words with strings.Fields
	words := strings.Fields(line)

	// empty string for generate result
	result := ""
	// invert the word order in the slice
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i] + " "
	}

	// remove extra space and return result
	return strings.TrimSpace(result)
}

func main() {
	sourseStr := "Hello Max Bye"
	reverseStr := reverseWords(sourseStr)

	fmt.Println("Sourse:", sourseStr)
	fmt.Println("After :", reverseStr)
}
