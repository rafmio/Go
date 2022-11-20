package main
import "fmt"

func main() {
    var arr [4] string
    arr[0] = "Kissy"
    arr[1] = "Missy"
    arr[2] = "Huggy"
    arr[3] = "Wuggy"
    
    fmt.Println("arr[0] = ", arr[0], "arr[1] = ", arr[1])
    fmt.Println(arr)
    
    primes := [6]int{10, 20, 30, 40, 50, 60}
    fmt.Println("primes: ", primes)
    
    for i := 0; i < len(arr); i++ {
        fmt.Printf("arr[%d] = %s\n", i, arr[i])
    }
}


// Arrays
// https://go.dev/tour/moretypes/6
