package main

import (
	"fmt"
)

type Animal struct {
	name string
}

func (a Animal) Run() {
	fmt.Println(a.name, "is running...")
}

func (a *Animal) RunFaster() {
	fmt.Println(a.name, "sometimes eat somebody...")
}

func main() {
	a := Animal{"Lion"}
	a.Run()
	a.RunFaster()
}

// Methods
// There are two types of receivers that are available in Go:
//      -- value receivers
//      -- pointer receivers
// https://golangdocs.com/methods-in-golang
