package main

import (
	"fmt"
)

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)

	integers := make([]int, 2)
	fmt.Println(integers)

	integers = nil
	fmt.Println(integers)

	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]

	fmt.Println(anArray)
	fmt.Println(refAnArray)

	anArray[4] = -100
	fmt.Println(refAnArray)

	s := make([]byte, 5)
	fmt.Println(s)
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	fmt.Println()
}
