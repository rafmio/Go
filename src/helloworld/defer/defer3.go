package main

import "fmt"

func main() {
    i := 0
    defer fmt.Println(i)
    i++
    fmt.Println(i)
    i++
    fmt.Println(i)
    
    for x := 0; x < 4; x++ {
        defer fmt.Printf("%d ", x)
    }
}
