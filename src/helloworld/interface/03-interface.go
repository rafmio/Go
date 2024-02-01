package main

import (
	"fmt"
)

type animal interface {
	makeSound()
}

type cat struct{}
type dog struct{}
type bird struct{}

func (c *cat) makeSound() {
	fmt.Println("meow!")
}

func (d *dog) makeSound() {
	fmt.Println("woof!")
}

func (b *bird) makeSound() {
	fmt.Println("tweet!")
}

func main() {
	var c, d, b animal = &cat{}, &dog{}, &bird{}
	c.makeSound()
	d.makeSound()
	b.makeSound()

	fmt.Printf("Type of c: %T, type of d: %T, type of b: %T\n", c, d, b)
}
