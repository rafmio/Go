package main
import "fmt"

// for Expample 2 --------------------------------------
func newCounter() func() int {
    GFG2 := 0
    return func() int {
        GFG2 += 1
        return GFG2
    }
}

func main() {
    // Example 1-----------------------------------------
    GFG := 0 
    
    counter := func() int {
        GFG += 1
        return GFG
    }
    
    fmt.Println("counter(): ", counter(), " GFG: ", GFG)
    fmt.Println("counter(): ", counter(), " GFG: ", GFG)
    fmt.Println("counter(): ", counter(), " GFG: ", GFG)
    fmt.Println("counter : ", counter)

    // Example 2-----------------------------------------

    counter2 := newCounter()
    
    fmt.Println(counter2())
    fmt.Println(counter2())
    fmt.Println(counter2)
    
    
}


// Function closures
// https://www.geeksforgeeks.org/closures-in-golang/
