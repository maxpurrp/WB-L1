package main

import (
	"fmt"
)

// stringRevers flips the string, considering Unicode chars
func stringRevers(s string) string {
	// convert string into a slice of rune to work correctly with Unicode
	runes := []rune(s)

	// get the len of rune slice
	length := len(runes)

	// start and end indexes a slice of runes
	start, end := 0, length-1

	// invert the order of runes in slice
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	// create and return new string from rune slice
	return string(runes)
}

func main() {
	// sourse string
	sourseStr := "GoDeveloperCool"
	// call func for reverse string
	reversStr := stringRevers(sourseStr)
	// print in stdout result
	fmt.Println("Исходная строка:", sourseStr)
	fmt.Println("Перевернутая строка:", reversStr)
}
