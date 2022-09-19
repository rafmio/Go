package main
import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("Kissy", "Missy")
    fmt.Println(a, b)
}

// Functions
// Multiple results
// https://go.dev/tour/basics/6
