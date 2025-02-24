package main
import (
    "bytes"
    "fmt"
)

func main() {
    // Example 1 -------------------------------------------------
    slc1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f', 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
    slc2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}
    slc3 := []byte{'%', 'g', 'e', 'e', 'k', 's', '%'}
    
    fmt.Println("Original Slice: ")
    fmt.Printf("slc1: %s\n", slc1) 
    fmt.Printf("slc2: %s\n", slc2)
    fmt.Printf("slc3: %s\n", slc3)
    
    res1 := bytes.Trim(slc1, "!#")
    res2 := bytes.Trim(slc2, "*^")
    res3 := bytes.Trim(slc3, "%")
    
    fmt.Println("Result: ")
    fmt.Printf("slc1: %s\n", res1)
    fmt.Printf("slc2: %s\n", res2)
    fmt.Printf("slc3: %s\n", res3)
    
    fmt.Println("-----------------------------------------------")
    // Example 2 -------------------------------------------------
    res4 := bytes.Trim([]byte("***W*elcome t*o my worl*d**"), "*")
    res5 := bytes.Trim([]byte("!@Fuck th!@e w!@ars!@!"), "!@")
    res6 := bytes.Trim([]byte("^^Tilli-milli-triamdey&&"), "^&")

    fmt.Println("Final slice after trim: ")
    fmt.Printf("res4: %s\n", res4)
    fmt.Printf("res5: %s\n", res5)
    fmt.Printf("res6: %s\n", res6)
}

// Trim a slice of bytes
// https://www.geeksforgeeks.org/how-to-trim-a-slice-of-bytes-in-golang/
