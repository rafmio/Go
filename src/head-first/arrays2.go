package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108

	var letters = [3]string{"a", "b", "c"}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
}
