package main
import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 15}
    
    var sls1 []int = primes[1:4]
    fmt.Println("sls1: ", sls1)
    
    primes[1] = 2000
    fmt.Println("primes: ", primes)
    fmt.Println("sls1: ", sls1)
    
    fmt.Println("sls1[0]: ", sls1[0])
    
    sls2 := []int{10, 20, 55, 66, 77, 88}
    fmt.Println("sls2: ", sls2)
    
    sls3 := sls2[2:5]
    fmt.Println("sls3: ", sls3)
    
    sls2[2] = 1000
    fmt.Println("sls2: ", sls2)
    fmt.Println("sls3: ", sls3)
}

// Slices
// https://go.dev/tour/moretypes/7
