package main

import (
	"fmt"
)

// search value in sorted slice and return index of value
// if he found, and -1, if not
func binSearch(slice []int, num int) int {
	// setting the initial range for search
	low, high := 0, len(slice)-1

	// The loop is executed until the lower bound exceeds the upper one.
	for low <= high {

		// calculating index of middle of array
		mid := (low + high) / 2

		//if slice[mid] is out num -> found and return
		if slice[mid] == num {
			return mid

		} else if slice[mid] < num { // if desired num is greater, shifting low to the right
			low = mid + 1
		} else { // if desired num is less, shifting high to the left
			high = mid - 1
		}
	}
	//return -1, if not found num in array
	return -1
}

func main() {
	//array must be sorted
	slice := []int{1, 4, 5, 6, 7, 8, 44, 55, 66, 88, 100}

	// use package sort for sorting slice
	// badArr := []int{4, 3, 2, 1, 6, 7, 4, 6}
	// sort.Ints(badArr)
	num := 7
	index := binSearch(slice, num)

	if index != -1 {
		fmt.Printf("Elememt %v was found by index %d\n", num, index)
	} else {
		fmt.Printf("Element %v not found in array", num)
	}
}
