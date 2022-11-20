package main

import "fmt"

func outer(name string) func() {
    text := "Modified " + name
    foo := func() {
        fmt.Println(text)
    }
    return foo
}

func main() {
    foo := outer("hello")
    foo()
}
