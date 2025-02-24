// https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/
// indirect recursion
package main

import (
	"fmt"
	"os"
)

func printOne(n int) {
	if n >= 0 {
		fmt.Println("in printOne():", n)
		printTwo(n - 1)
	}
}

func printTwo(n int) {
	if n >= 0 {
		fmt.Println("in printTwo():", n)
		printOne(n - 1)
	}
}

func main() {
	printOne(10)
	printOne(0)
	printOne(-1)

	os.Exit(0)
}
