package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Задаем значения переменных a и b, большие чем 2^20.
	a := new(big.Int).Lsh(big.NewInt(2), 23)
	b := new(big.Int).Lsh(big.NewInt(3), 31)

	// multiplication
	multiplicationResult := new(big.Int).Mul(a, b)
	fmt.Printf("Multiplication: %s * %s = %s\n", a.String(), b.String(), multiplicationResult.String())

	// division
	divisionResult := new(big.Int).Div(a, b)
	fmt.Printf("Division: %s / %s = %s\n", a.String(), b.String(), divisionResult.String())

	// addition
	additionResult := new(big.Int).Add(a, b)
	fmt.Printf("Addition: %s + %s = %s\n", a.String(), b.String(), additionResult.String())

	// subtraction
	subtractionResult := new(big.Int).Sub(a, b)
	fmt.Printf("Subtraction: %s - %s = %s\n", a.String(), b.String(), subtractionResult.String())
}
