package main

import "fmt"

func main() {
	var value int = 2
	var pointer *int = &value
	fmt.Println(value)
	fmt.Println(pointer)
	fmt.Println(*pointer)
}
