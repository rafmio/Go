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
    fmt.Println()
    
    const price2, tax2 float32 = 275, 27.50
    const quantity2, inStock2 = 2, true
    fmt.Println("Total:", 2 * quantity2 * (price2 + tax2))
    fmt.Println("In stock2: ", inStock2)
    fmt.Println()
    
    var price3 float32 = 275.00
    var tax3 float32 = 27.50
    fmt.Println(price3 + tax3)
    price3 = 300
    fmt.Println(price3 + tax3)
    fmt.Println()
    
    // Omitting var's type:
    var price4 = 275.00
    var price5 = price4
    fmt.Println("price4:", price4)
    fmt.Println("price5:", price5)
    fmt.Println()
    
    var price6, tax6 = 275.00, 27.50
    fmt.Println(price6 + tax6)
    fmt.Println()
    
    price7 := 275.00
    fmt.Println(price7)
    fmt.Println()
    
    price8, tax8, inStock8 := 275.00, 27.50, true
    fmt.Println("Total:", price8 + tax8)
    fmt.Println("In stock:", inStock8)
//     var price8, tax8 = 200.00, 25.00
//     fmt.Println("Total 2:", price8 + tax8)
    price9, tax8 := 200.00, 25.00
    fmt.Println("Total 2:", price9 + tax8)
    fmt.Println()
    
    // Blank identifier
    price10, tax10, inStock10, _ := 275.00, 27.50, true, true
    var _ = "Alice"
    fmt.Println("Total:", price10 + tax10)
    fmt.Println("In stock:", inStock10)
}
