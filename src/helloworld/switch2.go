package main
import "fmt"

func main() {
    var value string = "five"
    switch value {
        case "one":
            fmt.Println("One")
        case "two", "three":
            fmt.Println("Two and Three")
        case "four", "five", "six":
            fmt.Println("Four, Five, Six")
    }
}


// Switch Statement in Go
// https://www.geeksforgeeks.org/switch-statement-in-go/
