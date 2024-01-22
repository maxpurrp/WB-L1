package main

import "fmt"

func appendIfNotExists(slice []int, num int) []int {
	for _, elem := range slice {
		if elem == num {
			return slice //elem in slice,do not add again
		}
	}
	return append(slice, num)
}

func getIntersection(firstSlice, secondSlice []int) []int {
	// create map fro unique value in firtsArr
	set := make(map[int]struct{})

	// create slice for saving intersections
	result := []int{}

	// add nums firstArr in map
	for _, num := range firstSlice {
		set[num] = struct{}{}
	}

	// check ever element from secondArr for presence in map
	for _, num := range secondSlice {
		if _, exists := set[num]; exists {
			// if num in two arrays - append(if not exists) in final map
			result = appendIfNotExists(result, num)
		}
	}
	return result
}

func main() {
	firstArr := []int{1, 4, 1, 1, 4}
	secondArr := []int{4, 1, 4, 1, 4}

	// call the func to find intersections
	intersections := getIntersection(firstArr, secondArr)

	// print result
	fmt.Println("Intersection:", intersections)
}
