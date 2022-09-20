package main
import "fmt"

func main() {
    // Example 1 ------------------------------------
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }

    fmt.Println(sum)

    fmt.Println("----------------------------------")
    // Example 2 ------------------------------------
    var sum2 int = 0
    for i := 0; i < 12; i++ {
        sum2 = (sum-1000 - (i * 2))
        fmt.Println("sum2 = ", sum2)
    }


}

// For
// Init and post statements are optional
// https://go.dev/tour/flowcontrol/2
