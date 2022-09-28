package main
import "fmt"

type student struct {
    int 
    string
    float64
}

func main() {
    value := student{123, "Bud", 903.22}
    
    fmt.Println("Enrollment number: ", value.int)
    fmt.Println("Student name:      ", value.string)
    fmt.Println("Package price:     ", value.string) 
}


// Anonymous structures
// https://www.geeksforgeeks.org/anonymous-structure-and-field-in-golang/
