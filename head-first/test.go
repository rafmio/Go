package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Stdin)
    stdinn := os.Stdin
    fmt.Printf("Type of stdinn = %T\n", stdinn)
}
