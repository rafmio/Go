package main
import "fmt"

func main() {
    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("slice q:  ", q)
    
    qq := []float32{12.33, 55.22, 42.14, 21.140}
    fmt.Println("slice qq: ", qq)
    
    rr := []bool{true, false, true, true, false, true}
    fmt.Println("slice rr: ", rr)
    
    ss := []struct {
        ii int
        bb bool
    }{
        {2, true},
        {3, false},
        {5, true}, 
        {11, false},
        {13, true},
    }
    
    fmt.Println(ss)
    fmt.Printf("Type of ss: %T\n", ss)
    
    fmt.Printf("Type of ss[0]: %T\n", ss[0])
    fmt.Println(ss[0], ss[1])
}


// Slice literal
// A slice literal is like an array literal without the length
// This is an array literal:
// [3]bool{true, true, false}
// And this is creates the same array as above, then
// builds a slice that references it:
// []bool{true, true, false}
// https://go.dev/tour/moretypes/9
