package main

import (
	"fmt"
	"time"
)

func writer(ch chan<- int) {
	//"infinite" loop for writing data(i) in channel
	for i := 1; ; i++ {
		ch <- i
		fmt.Printf("Write in channel : %v \n", i)
	}
}

func reader(ch <-chan int) {
	//loop for reading data from channel and print data in stdout
	for {
		data := <-ch
		fmt.Printf("read from channel : %v \n", data)
	}
}

func main() {
	// creating channel for sharing data
	dataChannel := make(chan int)

	// setting the working time
	timeDuration := 2

	// Launching goroutine to read from channel
	go reader(dataChannel)

	// Launching goroutine to write in channel
	go writer(dataChannel)

	// Waiting time duration
	time.Sleep(time.Duration(timeDuration) * time.Second)
	fmt.Println("Program done")
}
