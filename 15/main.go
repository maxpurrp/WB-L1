package main

import (
	"fmt"
)

var justString string

// func create string and return string of the specified size
func createHugeString(size int) string {
	// creating slice of a given size
	hugeString := make([]byte, size)

	// convert the slice to string
	return string(hugeString)
}

func someFunc() {
	size := 100
	v := createHugeString(1 << 10)
	if len(v) >= size {
		justString = v[:size]
	} else {
		fmt.Println("string is too small")
	}

}

func main() {
	someFunc()

}
