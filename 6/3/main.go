package main

import (
	"fmt"
	"time"
)

func work() {
	// goroutine starts
	fmt.Println("Goroutine starts")
	//something else
	fmt.Println("Goroutine stoped")
	//return stops the function and the goroutine
	return

}

func main() {
	go work()
	// Waiting before goroutine finish
	time.Sleep(2 * time.Second)
	fmt.Println("Основная горутина завершила выполнение")
}
