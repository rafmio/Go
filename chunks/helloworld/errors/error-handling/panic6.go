package main

import "fmt"

func main() {
    names := []string{
        "lobster",
        "sea urchin",
        "sea cucumber",
    }
    fmt.Println("My favorite sea creature is: ", names[len(names)])
}

// panic: runtime error: index out of range [3] with length 3
