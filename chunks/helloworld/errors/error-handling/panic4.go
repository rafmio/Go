package main

import "os"

func main() {
    panic("a problem")
    _, err := os.Create("~/go/src/helloworld/file-panic.txt")
    if err != nil {
        panic(err)
    }
}


// https://gobyexample.com/panic
