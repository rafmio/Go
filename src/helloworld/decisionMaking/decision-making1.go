package main
import "fmt"

func main() {
    var v1 int = 700
    if v1 == 100 {
        fmt.Printf("Value of v1 is 100\n")
    } else if v1 == 200 {
        fmt.Printf("Value of v1 is 200\n")
    } else if v1 == 300 {
        fmt.Printf("Value of v1 is 300")
    } else {
        fmt.Printf("None of the values is matching\n")
    }
}
