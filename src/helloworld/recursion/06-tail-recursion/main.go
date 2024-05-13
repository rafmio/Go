// https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/
// tail recursion
package main

import (
	"fmt"
)

func printNum(n int) {
	if n > 0 {
		fmt.Println("printNum():", n)
		printNum(n - 1)
	}
}

func main() {
	printNum(5)
}
