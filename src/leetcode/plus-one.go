// You are given a large integer represented as an integer array digits,
// where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant in left-to-right
// order. The large integer does not contain any leading 0's.

// Increment the large integer by one and return the resulting array of digits.

package main

import (
	"fmt"
)

func main() {
	digits := []int{4, 3, 2, 1}
	incrementedDigits := plusOne(digits)
	fmt.Println(incrementedDigits)
}

func plusOne(digits []int) []int {
	lenDigits := len(digits)
	fmt.Println("lenDigits: ", lenDigits)

	for i := 0; i < lenDigits; i++ {
		fmt.Println(digits[i], " ")
	}

	for i := (lenDigits - 1); i <= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			fmt.Printf("digits[%d] = %d\n", i, digits[i])
			digits[i]++
			// break
		}
		fmt.Println("current i: ", i)
	}

	return digits
}
