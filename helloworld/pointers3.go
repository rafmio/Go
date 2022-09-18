package main
import "fmt"

func main() {
    // Example 1 -------------------------------------
    x := 0xFF
    y := 0x9C
    
    fmt.Printf("Type of variable x is: %T\n", x)
    fmt.Printf("Value of x in hexadecimal is: %X\n", x)
    fmt.Printf("Value of x in decimal is:     %v\n", x)
    
    fmt.Printf("Type of variable y is: %T\n", y)
    fmt.Printf("Value of y in hexadecimal is: %X\n", y)
    fmt.Printf("Value of y in decimal is:     %v\n", y)
    
    fmt.Println("-------------------------------------")
    // Example 2 -------------------------------------
    var x2 int = 5748
    var p *int
    p = &x2 
    
    fmt.Println("Value stored in x2 =         ", x2)
    fmt.Println("Addres of x =                ", &x)
    fmt.Println("Value stored in variable p = ", p)
    fmt.Println("Dereferencing p (*p):        ", *p)
    
    fmt.Println("-------------------------------------")
    // Example 3 -------------------------------------
    var x3 *int
    fmt.Println("x3 = ", x3)
    
    fmt.Println("-------------------------------------")
    // Example 4 -------------------------------------
    var x4 = 567
    var p1 = &x4
    
    fmt.Println("x4  =  ", x4)
    fmt.Println("&x4 =  ", &x4)
    fmt.Println("p1  =  ", p1)
    fmt.Println("*p  =  ", *p1)
    
     fmt.Println("-------------------------------------")
    // Example 5 -------------------------------------
}


// Pointers 
// https://www.geeksforgeeks.org/pointers-in-golang/
