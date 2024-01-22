package main

import (
	"fmt"
	"sync"
)

// calculation is a function that performs the calculation
// of the square of a number and prints the result.
// It uses sync.WaitGroup to wait for the completion of all goroutines.
func calculation(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	result := num * num
	fmt.Printf("result of calculation is : %v \n", result)
}

func main() {
	// used to wait for the completion of all goroutines
	var wg sync.WaitGroup
	numsLst := [5]int{2, 4, 5, 6, 10}
	for i := 0; i < len(numsLst); i++ {
		//Increment the goroutine counter before starting a new goroutine
		wg.Add(1)
		go calculation(&wg, numsLst[i])
	}
	// Wait for the completion of all goroutines
	wg.Wait()
}
