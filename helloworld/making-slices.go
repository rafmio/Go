package main
import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a)
    
    for i := 0; i < cap(a); i++ {
        a[i] = (i + 1) * i
    }
    printSlice("a", a)
    
    b := make([]int, 10, 10)
    printSlice("b", b)
    
    for i := 0; i < len(b); i++ {
        b[i] = (i + 1) * i
    }
    printSlice("b", b)
    
    c := b[:2]
    printSlice("c", c)
    
    d := c[2:5]
    printSlice("d", d)
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len = %d cap = %d %v\n", 
               s, len(x), cap(x), x)
}

// Slices can be created with the built-in 'make' function;
// this is how you create dynamically-sized arrays
// The 'make' funcions allocates a zeroed array and returns 
// a slice that refers to that array

// To specify a capacity, pass a third argument to 'make':
// b := make([]int, 0, 5)   - len(b)=0, cap(b)=5
// b = b[:cap(b)]           - len(b)=5, cap(b)=5
// b = b[1:]                - len(b)=5, cap(b)=4
