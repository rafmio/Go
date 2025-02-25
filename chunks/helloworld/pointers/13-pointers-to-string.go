package main

import (
	"fmt"
)

func printString(str *string) {
	if str == nil {
		fmt.Println("String is nil")
		return
	}
	fmt.Println(*str)
}

func main() {
	var str *string
	printString(str)

	str = new(string)
	*str = "Hello, World!"
	printString(str)

	str = nil
	printString(str)
}
