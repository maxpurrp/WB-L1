package main

import (
	"fmt"
	"time"
)

// native delay function without using time.Sleep
func mySleep(seconds int) {
	// create channel for waiting time duration
	ch := time.After(time.Duration(seconds) * time.Second)
	//goroutine will be blocked, intill the value is receieved frim channel
	<-ch
}

func getTime() (int, int, int) {
	hour, min, sec := time.Now().Clock()
	return hour, min, sec
}
func main() {
	//get current time
	hour, min, sec := getTime()
	fmt.Printf("Start on %02d:%02d:%02d\n", hour, min, sec)

	// call func for sleep
	mySleep(5)

	//get current time after sleep
	hour, min, sec = getTime()
	fmt.Printf("End on   %02d:%02d:%02d\n", hour, min, sec)
}
