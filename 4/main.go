package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {

	var numWorkers int
	fmt.Print("Write count of workers: ")
	fmt.Scan(&numWorkers)
	// data exchange channel for exchanging data
	// between the main goroutine and worker goroutines.
	dataChannel := make(chan string)

	// Channel for handling Ctrl+C interrupts
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT)

	// anonymous goroutine to handle Ctrl+C and close dataChannel.
	go func() {
		<-signalChannel
		fmt.Println("Closing the program")
		close(dataChannel)
	}()

	// starts workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel)
	}
	// main goroutine for writing data in channel
	go func() {
		defer close(dataChannel)
		for {
			var data string
			fmt.Println("Enter the data to write in channel")
			fmt.Scan(&data)
			dataChannel <- data
		}
	}()
	// Waiting before all goroutine finish
	wg.Wait()
	fmt.Println("Main goroutine has ended")
}

// func representing a worker goroutine that processes data from the channel.
func worker(id int, dataChannel <-chan string) {
	defer wg.Done()
	for data := range dataChannel {
		fmt.Printf("Worker %d: %s\n", id, data)
	}
}
