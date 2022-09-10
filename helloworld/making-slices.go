package main
import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a) 
    
    b := make([]int, 0, 5)
    printSlice("b", b)
    
    c := b[:2]
    printSlice("c", c)
    
    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}

// Slices cat be cereated with the built-in 'make' funcion; this
// is how you create dynamically-sized arrays
// The 'make' funcion allocates a zeroed array and returns a slice 
// that reffers to that array 
