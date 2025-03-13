package main

import (
    "fmt"
    "errors"
)

func main() {
    err := errors.New("emit macho dwarf: elf header corrupted")
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println()
    fmt.Printf("Type of err: %T\n", err)
}


// https://pkg.go.dev/errors
