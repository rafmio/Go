package main
import "fmt"

type Employee struct {
    name    string
    empid   int
}

func main() {
    // Example 1-------------------------------------
    emp := Employee{"Justin Jonhson", 19078}
    pts := &emp 
    
    fmt.Println("Direct accessing: ")
    fmt.Println("emp.name:  ", emp.name)
    fmt.Println("emp.empid: ", emp.empid)
    
    fmt.Println("Accessing via pointer: ")
    fmt.Println(pts)
    fmt.Println(pts.name)
    fmt.Println((*pts).name)
    fmt.Println(pts.empid)
    fmt.Println((*pts).empid)
    
    fmt.Println("-----------------------------------")
    // Example 2-------------------------------------
    pts.name = "Jerry Lee Luis"
    fmt.Println("After modifying pts.name = '...' :", pts)
}

// Pointers to a struct 
// https://www.geeksforgeeks.org/pointer-to-a-struct-in-golang/
