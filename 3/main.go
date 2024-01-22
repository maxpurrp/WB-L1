package main

import (
	"fmt"
	"sync"
)

var sum int

// calculation is a function that performs the calculation of the square of a number
// and increments the global sum variable with the result.
// It uses sync.WaitGroup to wait for the completion of all goroutines.
func calculation(wg *sync.WaitGroup, num int) {
	// decrement the goroutine counter upon completion
	defer wg.Done()
	result := num * num
	sum += result
}

func main() {
	// used to wait for the completion of all goroutines
	var wg sync.WaitGroup
	numsLst := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < len(numsLst); i++ {
		// increment the goroutine counter before starting a new goroutine
		wg.Add(1)
		go calculation(&wg, numsLst[i])

	}
	// wait for the completion of all goroutines
	wg.Wait()
	fmt.Println("result is :", sum)
}
