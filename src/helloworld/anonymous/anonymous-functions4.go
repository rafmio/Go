package main

import "fmt"

func outer(name string) {
    text := "Modified " + name 
    foo := func() {
        fmt.Println(text)
    }
    foo()
}

func main() {
    outer("hello")
}

// https://kovardin.ru/articles/go/zamykaniya/
