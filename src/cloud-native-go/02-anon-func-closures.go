// Anonymous functions and closures
package main

import (
	"fmt"
)

func sum(x, y int) int {
	return x + y
}

func product(x, y int) int {
	return x * y
}

func main() {
	var f func(int, int) int

	f = sum
	fmt.Println(f(3, 5))

	f = product
	fmt.Println(f(3, 5))
}
