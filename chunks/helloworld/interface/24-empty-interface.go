// https://go.dev/tour/methods/14
package main

import (
	"fmt"
	"os"
)

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = map[string]int{"key": 42}
	describe(i)

	i = []int{1, 2, 3}
	describe(i)

	i = func(x, y int) int { return x + y }
	describe(i)

	os.Exit(0)
}
