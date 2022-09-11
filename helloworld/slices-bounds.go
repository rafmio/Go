package main
import "fmt"

func main() {
    s := []int{12, 32, 434, 6313, 121, 4, 5, 353, 51}
    s = s[1:7]
    fmt.Println(s)
    
    s = s[:4]
    fmt.Println(s)
    
    s = s[1:]
    fmt.Println(s)
}


// When slicing, you may omit the high or low bounds to use their
// defaults instead
// The defaults is zero for the low bound and the length of the 
// slice for the high bound
// My notes (MAY BE WRONG!!!): in this example s redefined by // 
// assinging new values of new slices
