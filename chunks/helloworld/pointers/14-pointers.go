package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func newPerson() *person {
	return &person{
		Name: "John Doe",
		Age:  30,
	}
}

func main() {
	p := newPerson()
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
