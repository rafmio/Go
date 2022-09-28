package main
import (
    "fmt"
    "sort"
)

func main() {
    // Example 1-----------------------------------------
    slc1 := []int{400, 600, 4, 696, 431, 935, 90, 99}
    slc2 := []int{-23, 5467, -65, 141, 69, -43, 100, -4}
    
    fmt.Println("Slices before: ")
    fmt.Println("slc1: ", slc1)
    fmt.Println("slc2: ", slc2)
    
    sort.Ints(slc1)
    sort.Ints(slc2)
    
    fmt.Println()
    fmt.Println("Slices after: ")
    fmt.Println("slc1: ", slc1)
    fmt.Println("slc2: ", slc2)

    fmt.Println("---------------------------------------")
    // Example 2-----------------------------------------
    slc3 := []int{12, 32, 43, 54, 65, 76, 87, 100, 111}
    slc4 := []int{-43, 100, -1, 99, 123, -12, 53, 1010, -2, 432}
    
    fmt.Println("Slices: ")
    fmt.Println("slc3: ", slc3)
    fmt.Println("slc4: ", slc4) 
    
    res3 := sort.IntsAreSorted(slc3)
    res4 := sort.IntsAreSorted(slc4)
    
    fmt.Println("Result: ")
    fmt.Println("Is slc3 sorted? : ", res3)
    fmt.Println("Is slc4 sorted? : ", res4)

    fmt.Println("---------------------------------------")
    // Example 3-----------------------------------------
    
}

// Sort slices
// https://www.geeksforgeeks.org/how-to-sort-a-slice-of-ints-in-golang/
