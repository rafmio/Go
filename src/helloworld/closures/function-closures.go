package main

import "fmt"

func adder() func(int) int {
    sum := 0
    
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    
    for i := 0; i < 10; i++ {
        fmt.Printf("pos = %d, neg = %d\n", pos(i), neg(-2 * i))
    }
}





// Function closures
// https://go.dev/tour/moretypes/25
