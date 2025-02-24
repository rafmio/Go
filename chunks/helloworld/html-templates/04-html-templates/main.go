package main

import (
	"fmt"
)

func sum(a int, b int) {
	fmt.Printf("Sum of %d and %d is: %d\n", a, b, a+b)
}

func main() {
	anonFunc := func(a int, b int) {
		fmt.Printf("Anonymous function: Sum of %d and %d is: %d\n", a, b, a+b)
	}

	sum(3, 5)
	anonFunc(7, 9)

	fmt.Printf("Type of sum(): %T\n", sum)
	fmt.Printf("Type of anonFunc(): %T\n", anonFunc)

	funcSls := make([]func(int, int), 3)
	funcSls[0] = sum
	funcSls[1] = anonFunc
	funcSls[2] = func(a int, b int) {
		fmt.Printf("Diff: %d - %d = %d\n", a, b, a-b)
	}

	for i, f := range funcSls {
		fmt.Printf("Address of %T is: at index %d is: %p\n", f, i, &f)
	}

	fmt.Printf("Address of funcSls: %p\n", funcSls)
}
