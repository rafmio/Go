package main

import "fmt"

func main() {
    var sss []int
    fmt.Println(sss, len(sss), cap(sss))
    if sss == nil {
        fmt.Println("nil!")
    }
}

// Nil slices
// https://go.dev/tour/moretypes/12
