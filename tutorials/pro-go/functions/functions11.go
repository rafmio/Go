// Returning multiple function results
package main

import "fmt"

func swapValues(first, second int) (int, int) {
	return second, first
}

func main() {
	val1, val2 := 10, 20
	fmt.Println("Before calling function:", val1, val2)
	val1, val2 = swapValues(val1, val2)
	fmt.Println("After calling function:", val1, val2)
}
