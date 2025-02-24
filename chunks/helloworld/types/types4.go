package main
import (
    "fmt"
    "sort"
)

func main() {
    var array [3]string
    array[1] = "be-be"
    fmt.Println(array)
    fmt.Println(len(array))
    fmt.Println("---1---")
    
    var array2 = [3]int{12, 34, 56}
    for _, i := range array2 {
        fmt.Println(i)
    }
    
    fmt.Println("---2---")
    
    var slice1 = []int{332, 118, 2, 4, 32, 64, 128, 256}
    for _, i := range slice1 {
        fmt.Println(i)
    }
    
    fmt.Println("---3---")

    b := slice1[2:6]
    fmt.Println(b)
    
    fmt.Println("---4---")

    sort.Ints(slice1)
    fmt.Println(slice1)
    
    fmt.Println("---5---")

    
    slice1 = append(slice1, 321, 555)
    fmt.Println(slice1)
    
    fmt.Println("---6---")
    
    slice2 := make([]int, 10)
    fmt.Println(slice2)
}


// Array, slices
