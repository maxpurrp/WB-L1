package main

import (
	"fmt"
)

func groupTemperatures(temps []float64, step float64) map[int][]float64 {
	// initialize the map for store average temp and values
	group := make(map[int][]float64)

	for i := 0; i < len(temps); i++ {
		// calculation average temp for key
		roundedTemp := int(temps[i]/step) * int(step)
		// save in map avg temp ang append in array value
		group[roundedTemp] = append(group[roundedTemp], temps[i])
	}

	return group
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0
	// calculation group for temperatures
	tempGrups := groupTemperatures(temps, step)
	for k, v := range tempGrups {
		fmt.Printf("%v : %v \n", k, v)
	}
}
