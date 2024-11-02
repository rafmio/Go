package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Positive: %d\n", pos(i))
		} else {
			fmt.Printf("Negative: %d\n", neg(-i))
		}
	}

	fmt.Printf("\nAlt:\n")
	for i := 0; i < 10; i++ {
		fmt.Printf("pos = %d, neg = %d\n", pos(i), neg(-2*i))
	}
}
