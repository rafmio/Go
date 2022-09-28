package main
import "fmt"

type Bird struct {
    name string
}

func (b Bird)Fly() {            // Declare method
    fmt.Println(b.name, "is flying...")
}

func main() {
    b := Bird{"Raven"}
    b.Fly()                     // Call method
}


// Golang methods
// A methods are just like functions, except it has special 
// argument and that is a receiver
// Syntax: 
// func(receiver receiverType) funcName(arg argType) returnType{}
// https://golangdocs.com/methods-in-golang
