package main

import (
	"fmt"
	"time"
)

// structure representing action of Human
type Action struct{}

func (a *Action) Walk() {
	fmt.Printf("I walking and sing song")
}

func (a *Action) Sleep() {
	fmt.Println("I sleep and see nightmares")
	time.Sleep(2 * time.Second)
}

// structure representing a person with actions, name, age, and height.
type Human struct {
	Action
	Name   string
	Age    int
	Height int
}

func main() {
	person := Human{
		Name:   "Max",
		Age:    25,
		Height: 195,
	}
	fmt.Printf("hello, %v \n", person.Name)
	person.Sleep()
	fmt.Printf("And my height is : %v", person.Height)
}
