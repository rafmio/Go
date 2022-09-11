package main
import "fmt"

func main() {
    array := []int{22, 33, 43, 54, 57, 89}
    fmt.Printf("array []: len = %d cap = %d\n", len(array), cap(array))
    
    // Slice the slice to give it zero length
    array = array[:0]
    printSlice(array)
    
    // Extend its length
    array = array[:4]
    printSlice(array)
    
    // Drop its first two values
    array = array[2:]
    printSlice(array)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// A slice has both a length and a capacity
// The length of slice if the number of elements it contains
// The capacity of a slice is the number of elements in the 
// underlying array, counting from the first element is the slice
