package main

import (
	"fmt"
)

// Person - structure,representing a common interface
// for working with information about a Human
type Person interface {
	GetAge() int
	GetName() string
	GetWeight() float32
	IsSmoking() bool
	GetBloodSystem() chan string
}

// Human - structure from 14st task
type Human struct {
	Age         int
	Name        string
	Weight      float32
	Smoking     bool
	BloodSystem chan string
}

// Implementation of the methods of the Person interface for the Human structure
func (m *Human) GetAge() int {
	return m.Age
}

func (m *Human) GetName() string {
	return m.Name
}

func (m *Human) GetWeight() float32 {
	return m.Weight
}

func (m *Human) IsSmoking() bool {
	return m.Smoking
}

func (m *Human) GetBloodSystem() chan string {
	return m.BloodSystem
}

func main() {
	me := &Human{
		Age:         25,
		Name:        "Max",
		Weight:      95,
		Smoking:     false,
		BloodSystem: make(chan string),
	}

	// use the Me object through the Person interface.
	PrintPersonInfo(me)
}

// Print Person Info accepts an object implementing the Person interface
// and outputs its information.
func PrintPersonInfo(p Person) {
	fmt.Printf("Name: %s, Age: %d, Weight: %.2f, Smoking: %t\n", p.GetName(), p.GetAge(), p.GetWeight(), p.IsSmoking())
}
