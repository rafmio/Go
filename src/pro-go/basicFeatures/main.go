package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("Value:", rand.Int())
    fmt.Println("Hello, Go")
    fmt.Println(20 + 20)
    fmt.Println(20 + 30)
    fmt.Println()
    
    // Typed constant:
    const price float32 = 275.00
    const tax float32 = 27.50
    fmt.Println("price + tax: ", price + tax)
    fmt.Println()
    
    // Untyped constat (type was ommitted):
    const quantity = 2
    fmt.Println("Total:", quantity * (price + tax))
}
