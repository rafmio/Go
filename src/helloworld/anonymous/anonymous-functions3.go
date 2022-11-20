package main

import "fmt"

func printMessage(message string) {
    fmt.Println(message)
}

func getPrintMessage() func(string) {
    return func(message string) {
        fmt.Println(message)
    }
}

func main() {
    printMessage("Hello function!")
    
    func(message string) {
        fmt.Println(message)
    } ("Hello anonymous function!")
    
    printfunc := getPrintMessage()
    printfunc("Hello anonymous function using caller")
}

// https://kovardin.ru/articles/go/zamykaniya/
