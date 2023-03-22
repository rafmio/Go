package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)

	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)

	fmt.Println()
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("cap: %d, len: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("cap: %d, len: %d\n", cap(aSlice), len(aSlice))
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}

	fmt.Println()
}
