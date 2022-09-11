package main
import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}

// Range
// The 'range' from of the 'for' loop iterals over a slice or map
// When ranging over a slice, two values are returned for each
// iteration. The first is the indexx, and the second is a copy of
// the element at that index 

