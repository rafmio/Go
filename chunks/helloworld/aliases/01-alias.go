package main

import (
	"fmt"
	"strings"
)

type MyInt = int
type stringProcessor func(string) string

func main() {
	var x MyInt = 10
	fmt.Printf("The type of x is %T\n", x) // The type of x is main.MyInt

	var processor stringProcessor
	processor = func(s string) string {
		return strings.ToUpper(s)
	}
	fmt.Println(processor("hello")) // HELLO
}
