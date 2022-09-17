package main
import (
    "bytes"
    "fmt"
)

func main() {
    // Example 1 -------------------------------------------------
    slc1 := []byte{'!', '!', 'G', 'o', 'g', 'y', 'f', 'r', 'o', 'm', 'G', 'e', 'o', 'r', 'g', 'i', 'a'}
    slc2 := []byte{'A', 'l', 'm', 'a'}
    slc3 := []byte{'%', 'g', '%', 'e', '%', 'v', '%', 'o', '%', 'r', '%', 'g'}

    fmt.Println("Original slices: ")
    fmt.Printf("slc1: %s", slc1)
    fmt.Printf("slc2: %s", slc2)
    fmt.Printf("slc3: %s", slc3)
    
    res1 := bytes.Split(slc1, []byte("!!"))
    res2 := bytes.Split(slc2, []byte(""))
    res3 := bytes.Split(slc3, []byte("%"))
    
    fmt.Println("\nAfter splitting: ")
    fmt.Printf("slc1: %s\n", res1)
    fmt.Printf("slc2: %s\n", res2)
    fmt.Printf("slc3: %s\n", res3)  

    fmt.Println("-----------------------------------------------")
    // Example 2 -------------------------------------------------
    res4 := bytes.Split([]byte("***Welcome , to, the, world, of, MasterBoy***"), []byte(","))
    res5 := bytes.Split([]byte("Now x you x are x able x to x fly x up"), []byte("x"))
    res6 := bytes.Split([]byte("Queen, Elizabeth"), []byte(""))
    res7 := bytes.Split([]byte(""), []byte(","))
    
    fmt.Println("Final values: ")
    fmt.Printf("res4: %s\n", res4)
    fmt.Printf("res5: %s\n", res5)
    fmt.Printf("res6: %s\n", res6)
    fmt.Printf("res7: %s\n", res7)
    
}


// Split a slice of bytes
// https://www.geeksforgeeks.org/how-to-split-a-slice-of-bytes-in-golang/
