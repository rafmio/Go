// https://go101.org/article/type-embedding.html
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person
	works []string
}

func main() {
	var gaga = Singer{Person: Person{"Gaga", 30}, works: []string{"Pocker Face", "Paparazzi", "Bad Romance"}}
	gaga.PrintName()
	gaga.Name = "Lady Gaga"
	(&gaga).SetAge(31)
	(&gaga).PrintName()
	fmt.Println(gaga.Age)

	for key, val := range gaga.works {
		fmt.Println(key, val)
	}
}
