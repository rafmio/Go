package main

import "fmt"

func main() {
    var i interface{} = "hello"
    fmt.Println("i: ", i)
    
    fmt.Println()
    s := i.(string)
    fmt.Println(s)
    
    s, ok := i.(string)
    fmt.Println(s, ok)
    
    f, ok := i.(float64)
    fmt.Println(f, ok)
}

// Type assertion
// https://go.dev/tour/methods/15
