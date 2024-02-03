// https://golangbyexample.com/interface-in-golang/
package main

import (
	"fmt"
)

type animal interface {
	breathe()
	walk()
	showage()
}

type lion struct {
	name string
	age  int
}

func (l lion) breathe() {
	fmt.Println(l.name, "breathes")
}

func (l lion) walk() {
	fmt.Println(l.name, "walk")
}

func (l lion) showage() {
	fmt.Println(l.name, "age is:", l.age)
}

func main() {
	var a animal
	a = lion{name: "lion", age: 10}
	a.breathe()
	a.walk()
	a.showage()
	fmt.Printf("Type of a: %T\n", a)

	cat := lion{name: "cat", age: 9}
	cat.breathe()
	cat.walk()
	cat.showage()
}
