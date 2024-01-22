package main

import (
	"fmt"
)

// this method is easier to understand, but there may be
// problems with large numbers, in particular data type overflow
func swap(x, y int) (int, int) {
	x = x + y // 30
	y = x - y // 10
	x = x - y // 20
	return x, y
}

// this method has no overflow problems
// but requires knowledge of bitwise operations

// bite XOR
func bitSwap(x, y int) (int, int) {

	x = x ^ y // x contains the differences between x and y
	y = x ^ y // since x contains differnce, now y stores the sourse x
	x = x ^ y // now x has sourse y
	return x, y
}

func main() {
	x, y := 10, 20
	fmt.Println("Before swap: ", x, y)
	x, y = swap(x, y)
	fmt.Println("After  swap :", x, y)
}
