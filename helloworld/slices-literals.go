package main
import "fmt"

func main() {
    q := []int{20, 3, 50, 7, 110, 150}
    fmt.Println(q)
    for i := 0; i < len(q); i++ {
        fmt.Println("q: ", q[i])
    }
    
    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)
    for i := 0; i < len(r); i++ {
        fmt.Println("r: ", r[i])
    }
    
    s := []struct{
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true}, 
        {7, true},
        {11, false},
        {150, true},
    }
    
    fmt.Println(s)
}
// A slice literal is like an array literal without the length
// This is an array literal:
// [3]bool{true, true, false}
// And this is creates the same array as above, then
// builds a slice that references it:
// []bool{true, true, false}
