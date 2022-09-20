package main
import (
  "fmt"
  "math"
)

func main() {
    // Type conversions
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x * x + y * y))
    var z uint = uint(f)

    fmt.Println(x, y, z)

    var aa float64 = 12.55
    var bb int = int(aa)
    fmt.Println(aa)
    fmt.Printf("Type of bb is %T, value: %v\n", bb, bb)

    fmt.Println("-------------------------------------")
    // Type inference

    v := 42
    fmt.Printf("v is of type %T\n", v)
}


// Type conversions
// https://go.dev/tour/basics/13

// Type inference
// https://go.dev/tour/basics/14
