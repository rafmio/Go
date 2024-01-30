package main

import (
	"fmt"
)

type Bird struct {
	name string
}

func (b Bird) Fly() {
	fmt.Println(b.name, "is flying...")
}

func main() {
	a := Bird{name: "Eagle"}
	a.Fly()

	b := Bird{"Raven"}
	b.Fly()

}

// Golang methods
// A methods are just like functions, except it has special
// argument and that is a receiver
// Syntax:
// func(receiver receiverType) funcName(arg argType) returnType{}
// https://golangdocs.com/methods-in-golang
