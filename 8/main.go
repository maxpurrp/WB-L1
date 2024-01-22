package main

import (
	"fmt"
)

func changeBit(value int64, pos int, bitValue int) int64 {
	// create mask, which set i-n bite to 1
	mask := int64(1) << uint(pos)

	// set or clear the bit depends on bitValue
	if bitValue == 1 {
		// bites OR, set 1 where a least one of the bits is 1
		// 01001
		// 10010
		// 11011 <- result of bites OR
		value |= mask
	} else {
		// bites AND NOT sets bits to 0 where corresponding bits in the second operand are 1.
		// 10101110
		// 01011100
		// 00000001 <- result of bites AND NOT
		value &= ^mask
	}

	return value
}

func main() {
	var num int64 = 99
	var position int = 4
	var bite int = 1

	fmt.Printf("Source value: %d\n", num)

	num = changeBit(num, position, bite)
	fmt.Printf("Result is %v, change %v position to %v bite \n", num, position, bite)

}
