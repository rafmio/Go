// Using pointers as function parameters
package main

import "fmt"

func swapValues(first, second int) {
	fmt.Println("Before swap:", first, second)
	temp := first
	first = second
	second = temp
	fmt.Println("After swap:", first, second)
}

func main() {
	val1, val2 := 10, 20
	fmt.Println("Before calling function", val1, val2)
	swapValues(val1, val2)
	fmt.Println("After calling function", val1, val2)
}
