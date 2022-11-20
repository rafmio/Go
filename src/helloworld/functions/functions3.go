package main
import "fmt"

func swap(a, b *int) int {
    var o int
    o = *a
    *a = *b
    *b = o
    return o
}

func main() {
    var p int = 10
    var q int = 20
    fmt.Println("p = %d, q = %d", p, q)
    
    swap(&p, &q)
    fmt.Println("p = %d, q = %d", p, q)
}


// Functions
// https://www.geeksforgeeks.org/functions-in-go-language/
