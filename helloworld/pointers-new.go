package main
import "fmt"

func main() {
    ptri := new(int)
    *ptri = 67
    
    fmt.Println(" ptri: ", ptri)
    fmt.Println("*ptri: ", *ptri)
    
    size := new(int)
    fmt.Printf("Size value is %d, type is %T, address is %v\n", 
               *size, size, size)
    *size = 85
    fmt.Println("New size value is: ", size)
    
    b := 255
    a := &b
    fmt.Printf("Value of a is: %v, dereferenced value is: %d\n", 
               a, *a)
}

// The 'new' function in Go returns a pointer to a type
