package main

import "fmt"

func main() {
    first := 100
    var second *int = &first
    
    first++
    
    fmt.Println("First:", first)
    fmt.Println("Second:", second)
    fmt.Println("Second:", *second)
    fmt.Println()
    
    first2 := 1000
    second2 := &first2
    first++
    fmt.Println("First:", first2)
    fmt.Println("Second:", *second2)
    *second2++
    fmt.Println("Second:", *second2)
    var myNewPointer *int
    myNewPointer = second2
    *myNewPointer++
    fmt.Println("First:", first2)
    fmt.Println("Second:", *second2)
    fmt.Println()
    
    first3 := 1000
    var second3 *int
    fmt.Println(second3)
    second3 = &first3
    fmt.Println(second3)
    fmt.Println()
    
    // Pointer to a Pointer
    first4 := 555
    second4 := &first4
    third4 := &second4
    fmt.Println("first4:", first4)
    fmt.Println("*second4:", *second4)
    fmt.Println("**third4:", **third4)
}
