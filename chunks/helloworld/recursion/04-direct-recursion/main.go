// https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/
package main

import (
	"fmt"
)

func factorial_calc(number int) int {
	// this is the base condition
	// if number is 0 or 1 the function will return 1
	if number == 0 || number == 1 {
		return 1
	}

	if number < 0 {
		fmt.Printf("Invalid number: ")
		return -1
	}

	return number * factorial_calc(number-1)
}

func main() {
	answer1 := factorial_calc(0)
	fmt.Println("answer1:", answer1)
	answer2 := factorial_calc(5)
	fmt.Println("answer2:", answer2)
	answer3 := factorial_calc(-1)
	fmt.Println("answer3:", answer3)
	answer4 := factorial_calc(10)
	fmt.Println("answer4:", answer4)
}
