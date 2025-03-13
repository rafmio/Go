package main

import "fmt"

func main() {
	var i interface{}
	descibe(i)

	i = 42
	descibe(i)

	i = "hello"
	descibe(i)
}

func descibe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
