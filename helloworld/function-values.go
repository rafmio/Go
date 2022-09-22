package main

import (
    "fmt"
    "math"
)

func sum(aaa, bbb int) int {
    return aaa + bbb
}

func mult(fn func(int, int) int) int {
    return fn(10, 20)
}

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x * x + y* y)
    }
    
    fmt.Println(hypot(5, 12))
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
    
    aaa := 100
    bbb := 200
    fmt.Println("sum(aaa, bbb): ", sum(aaa, bbb))
    fmt.Println("mult(sum):     ", mult(sum))
    
}

// Function values
// Functions are values too. They can be passed around just like
// other values
// Function values may be used as function argument and return values
// https://go.dev/tour/moretypes/24
