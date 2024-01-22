package main

import (
	"fmt"
	"sync"
)

// wg is used for wait all goroutines
var wg sync.WaitGroup

// func, which send nums in channel
func reader(inpChan chan int) {
	//decrement count goroutines by one when goroutines completes
	defer wg.Done()
	//close channel upon exiting the function
	defer close(inpChan)
	//array of values to be sent into the channel
	numbers := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(numbers); i++ {
		//send in first channel value from array
		inpChan <- numbers[i]
	}
}

func calculation(inpChan chan int, outChan chan int) {
	//decrement count goroutines by one when goroutines completes
	defer wg.Done()
	//close channel upon exiting the function
	defer close(outChan)
	//iterate over the nums received from channel
	for num := range inpChan {
		//calculate value and send to second channel
		outChan <- num * 2
	}
}

func main() {
	// channels for transfer data
	inpChan := make(chan int)
	outChan := make(chan int)

	// goroutine for write nums in channel
	wg.Add(1)
	go reader(inpChan)

	// goroutine for calculation and writing nums in second channel
	wg.Add(1)
	go calculation(inpChan, outChan)

	// goroutine to wait for the completion of other goroutines
	go func() {
		wg.Wait()
	}()
	//iterate over result from second channel
	for result := range outChan {
		//print value in stdout
		fmt.Println(result)
	}
}
