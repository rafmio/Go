// Type assertion
// https://golangbyexample.com/interface-in-golang/
package main

import "fmt"

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

func main() {
	var a animal

	a = lion{age: 10}
	printA(a)
}

func printA(a animal) {
	l, ok := a.(lion)
	fmt.Println("Age:", l.age)
	fmt.Println("ok:", ok)

	if d, ok := a.(dog); ok {
		fmt.Println(d)
	} else {
		fmt.Println("This is not dog")
		fmt.Println("ok:", ok)
	}
}
