// https://golangbyexample.com/interface-in-golang/
package main

import "fmt"

type animal interface {
	breathe()
	walk()
	farts()
}

type human interface {
	animal
	speak()
}

type employee struct {
	name string
}

func (e employee) breathe() {
	fmt.Println("Employee breathes")
}

func (e employee) walk() {
	fmt.Println("Employee walk")
}

func (e employee) farts() {
	fmt.Println("Employee farts")
}

func (e employee) speak() {
	fmt.Println("Employee speaks")
}

func main() {
	var h human

	h = employee{name: "Jozeph"}
	h.breathe()
	h.walk()
	h.farts()
	h.speak()

	fmt.Println()
	fmt.Printf("Type of h is: %T\n", h)
}
