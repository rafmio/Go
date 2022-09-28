package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
    
    defer fmt.Println("Second defer funcion")
    
    fmt.Println("Kissy-Missy")
}


// Defer
// https://go.dev/tour/flowcontrol/12

// A defer statemens defers the execution of a function until
// the surrounding function returns
