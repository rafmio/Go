// https://gobyexample.com/struct-embedding
package main

import "fmt"

type base struct {
	name string
	num  int
}

func (b base) describe() string {
	return fmt.Sprintf("the %v's base with num=%v", b.name, b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			name: "cont",
			num:  300,
		},
		str: "some str",
	}

	fmt.Printf("co={name: %v, num: %v, str: %v}\n", co.name, co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
