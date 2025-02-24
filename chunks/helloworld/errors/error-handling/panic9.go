package main

import "fmt"

func main() {
    defer func () {
        fmt.Println("Hello the deffered funcion!")
    }()
    
    panic("oh no!")
}
