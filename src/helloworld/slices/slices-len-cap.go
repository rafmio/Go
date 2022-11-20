package main
import "fmt"

func main() {
    arr := [9]int{2, 4, 8, 16, 32, 64, 128, 256, 512}
    fmt.Println("arr: ", arr)
    fmt.Printf("len(arr) = %d, cap(arr) = %d\n", len(arr), cap(arr))
    
    sls1 := arr[2:8]
    fmt.Println()
    fmt.Println("sls1: ", sls1)
    fmt.Printf("len(sls1) = %d, cap(sls1) = %d\n", len(sls1), cap(sls1))
    
    sls1 = sls1[2: ]
    fmt.Println()
    fmt.Println("sls1: ", sls1)
    fmt.Printf("len(sls1) = %d, cap(sls1) = %d\n", len(sls1), cap(sls1))
    
    sls1 = sls1[1:4]
    fmt.Println()
    fmt.Println("sls1: ", sls1)
    fmt.Printf("len(sls1) = %d, cap(sls1) = %d\n", len(sls1), cap(sls1))
    
    sls1[0] = 1000
    fmt.Println()
    fmt.Println("arr: ", arr)
    fmt.Println("sls1: ", sls1)
}


// Slice length and capacity
// https://go.dev/tour/moretypes/11
