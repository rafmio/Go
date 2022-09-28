package main 
import "fmt"

func main() {
    i := 42
    f := func() {
        j := i / 2
        fmt.Println(j)
    }
    
    f()
}

// Function closure
// https://golangdocs.com/closures-in-golang
