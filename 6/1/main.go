package main

import (
	"fmt"
)

// exitFunc is a function representing a goroutine that performs
// a simple task and signals its completion.
// It takes a channel 'exit' as a parameter to signal when
// the goroutine has finished.
func exitFunc(exit chan bool) {
	fmt.Println("goroutine starts")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// signal that the goroutine has finished by sending a value to the 'exit' channel
	exit <- true
	fmt.Println("goroutine finish")

}

func main() {
	//create channel for flag
	exit := make(chan bool)

	//starts goroutine exitFunc
	go exitFunc(exit)

	//waiting for a signal from goroutine
	<-exit

	fmt.Println("Exit")
}
