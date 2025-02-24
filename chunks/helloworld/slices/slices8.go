package main

import "fmt"

func modSls(sls []int) {
    for i := range sls {
        sls[i]++
    }
}

func main() {
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    var sls []int = []int{6, 7, 8, 9, 10}
    fmt.Println(arr)
    fmt.Println(sls)
    fmt.Printf("Type of arr: %T, type of sls: %T\n", arr, sls)
    fmt.Println(len(arr), cap(arr))
    fmt.Println(len(sls), cap(sls))
    
    modSls(sls)
    fmt.Println(sls)
    modSls(sls)
    fmt.Println("sls:", sls)
    fmt.Println("arr:", arr)
    arr[0] = 300
    fmt.Println(arr)
}
