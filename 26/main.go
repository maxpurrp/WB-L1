package main

import "fmt"

func checkUnique(word string) bool {
	// use map for track uniaue chars
	seen := make(map[rune]bool)

	// iterate over each chars in word
	for _, char := range word {
		// if char has in map -> string has duplicates
		if seen[char] {
			return false
		}

		// mark current char as encounter
		seen[char] = true
	}

	// return true, if all chars encounter  unique
	return true
}

func main() {

	words := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, word := range words {
		fmt.Println(checkUnique(word))
	}
}
