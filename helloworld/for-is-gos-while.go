package main
import "fmt"

func main() {
    // Example 1 -----------------------
    sum := 1
    for sum < 1000 {
      sum += sum
    }
    fmt.Println(sum)

    // Example 2 -----------------------
    sum2 := 10
    for sum2 < 300 {
        sum2 += sum2
    }
    fmt.Println(sum2)

    // Example 3 -----------------------
    var sum3 int = 15
    for sum3 < 300 {
        sum3 += sum3
    }
    fmt.Println(sum3)

    // Example 4 -----------------------
    var sum4 = 50
    for sum4 > 25 {
        sum4 -= 5
        fmt.Printf("sum4 = %d\n", sum4)
    }
}

// For is Go's while
// https://go.dev/tour/flowcontrol/3
