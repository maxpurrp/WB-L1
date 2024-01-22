package main

import "fmt"

// simple structure for different types
type Me struct {
	Age         int
	Name        string
	Weight      float32
	Smoking     bool
	BloodSystem chan string
}

//get type of value with type assertion
func getType(value interface{}) {
	switch t := value.(type) {
	case int:
		fmt.Println("Type int", t)
	case float32:
		fmt.Println("Type float32", t)
	case string:
		fmt.Println("Type string", t)
	case bool:
		fmt.Println("Type bool", t)
	case chan string:
		fmt.Println("Type chan string")
	default:
		fmt.Printf("Unknown type: %v \n", t)
	}
}

func main() {
	//creating simple values of different types
	bloodChan := make(chan string)
	person := Me{
		Age:         25,
		Name:        "Max",
		Weight:      95.5,
		Smoking:     false,
		BloodSystem: bloodChan,
	}
	//call func for different fields of struct
	getType(person.Age)
	getType(person.Name)
	getType(person.Weight)
	getType(person.Smoking)
	getType(person.BloodSystem)

}
