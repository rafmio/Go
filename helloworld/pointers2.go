package main
import "fmt"

func main() {
    var x int = 5734
    var p *int
    p = &x
    fmt.Println("x = ", x)
    fmt.Println("Addr of x: ", &x)
    fmt.Println("Value stored in p: ", p)
    fmt.Println("Dereferensing p (*p): ", *p)
    
    // Zero-value of a pointer == nil
    var s *int
    fmt.Println("\nValue of s: ", s)
}


// Pointers
// The default value of zero-value of a pointer is always nil
