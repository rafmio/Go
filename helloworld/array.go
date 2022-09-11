package main
import "fmt"

func main() {
    var array [4]string
    array[0] = "Hello"
    array[1] = "Mello"
    array[2] = "Kissy"
    array[3] = "Missy"
    
    for i := 0; i < len(array); i++ {
        fmt.Println(array[i])
    }
    fmt.Println(array)
    
    primes := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
    for i := 0; i < len(primes); i++ {
        fmt.Println(primes[i])
    }
    fmt.Println(primes)
}

// The type [n]T is an array of n values of type T
// An array's length is part of its type, so arrays cannot be 
// resized
// But Go provides a convinient way of working with arrays
