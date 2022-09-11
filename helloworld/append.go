package main
import "fmt"

func main() {
    var s []int
    printSlice(s)
    
    // Append works on nil slices
    s = append(s, 0)
    printSlice(s)
    
    // The slice grows as needed
    s = append(s, 1, 58, 36)
    printSlice(s)
    
    // We can add more than one element at a time
    s = append(s, 12, 343, 453, 500)
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cat=%d %v\n", len(s), cap(s), s)
    fmt.Printf("Type of s[] is: %T\n\n", s)
}

// func append
// https://go.dev/tour/moretypes/15
// func append(slice []Type, elems ...Type) []Type
// https://pkg.go.dev/builtin#append
// The resulting value of append is a slice containing all the
// elements of the original slice plus the provided values

// If the backing array of 's' is to small to fit all the given
// values a bigger array will be allocated. The returned slice 
// will point to the newly allocated array
