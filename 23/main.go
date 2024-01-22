package main

import (
	"fmt"
)

func removeElem(array []int, index int) []int {
	// check valid target
	if index < 0 || index >= len(array) {
		fmt.Println("Incorrect index")
		return array
	}

	// create new slice without target index
	result := append(array[:index], array[index+1:]...)

	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	indRemove := 2
	fmt.Println("Slice before removing index:", numbers)
	numbers = removeElem(numbers, indRemove)

	fmt.Println("Slice after removing index: ", numbers)
}
