package main
import "fmt"

func main() {
    var sls []int
    fmt.Println(sls, len(sls), cap(sls))
    if sls == nil{
        fmt.Println("nil!")
    }
}


// The zero value of a slice is nil
// A nil has a length and capacity of 0 and has no underlying array
