package main
import "fmt"

func main() {
    q := []int{2, 3, 5, 7, 11, 15, 20, 25, 128}
    fmt.Println(q)
    
    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)
    
    s:= []struct{
        i int
        b bool
    }{
        {2, true},
        {3, false},
        {5, true},
        {7, true},
        {11, false},
        {15, true},
        {20, true},
        {25, false},
        {128, true},
    }
    fmt.Println(s)
}
