// https://www.geeksforgeeks.org/different-types-of-recursion-in-golang/
// head recursion
package main

import (
	"fmt"
)

func printNum(n int) {
	if n > 0 {
		printNum(n - 1)
		fmt.Println("printNum():", n)
	}
}

func main() {
	printNum(5)
}
