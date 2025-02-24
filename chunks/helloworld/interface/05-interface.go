package main

import (
	"fmt"
)

type animal interface {
	walker
	runner
}

type bird interface {
	walker
	flier
}

type flier interface {
	fly()
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type cat struct{}
type eagle struct{}
type mouse struct{}

func (c *cat) walk() {
	fmt.Println("cat is walking")
}

func (c *cat) run() {
	fmt.Println("cat is running")
}

func (e *eagle) walk() {
	fmt.Println("eagle is walking")
}

func (e *eagle) fly() {
	fmt.Println("eagle is flying")
}

func (m *mouse) walk() {
	fmt.Println("mouse is walking")
}

func (m *mouse) run() {
	fmt.Println("mouse is running")
}

func walk(w walker) {
	w.walk()
}

func main() {
	var c animal = &cat{}
	c.walk()
	c.run()

	var e bird = &eagle{}
	e.walk()
	e.fly()

	var m animal = &mouse{}
	m.walk()
	m.run()

	fmt.Println()
	walk(c)
	walk(e)
	walk(m)
}
