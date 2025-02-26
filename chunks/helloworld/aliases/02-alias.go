package main

import "fmt"

type mathOperation = func(int, int) int

func add(a, b int) int {
	return a + b
}

func main() {
	var operation mathOperation = add
	result := operation(10, 20)
	fmt.Println(result)
	fmt.Printf("Type of 'operation': %T\n", operation)
}
