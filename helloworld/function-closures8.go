package main

import "fmt"
 
func main() { 
    // Example 1 ------------------------------------------------
    l := 20
    b := 30
    
    func () {
        var area int
        area = l * b
        fmt.Println(area) 
        fmt.Printf("Type of area: %T\nType of l: %T\nType of b: %T\n", area, l, b)
    } ()
    
    fmt.Println("------------------------------------------------")
    // Example 2 ------------------------------------------------
    
    for i := 10.0; i < 100; i += 10.0 { 
        rad := func () float64 {
            return i * 39.370 
        } ()

        fmt.Printf("%.2f Meters = %.2f Inch\n", i, rad)
    }
} 


// Function closure
// https://www.golangprograms.com/closures-functions-in-golang.html
