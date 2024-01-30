// https://golangdocs.com/methods-in-golang
package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	gender string
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.name, "Age:", p.age, "Gender:", p.gender)
}

func main() {
	names := make([]string, 0)
	names = append(names, "Stephany")
	names = append(names, "Harry")
	names = append(names, "Angelina")
	names = append(names, "Gerge")
	names = append(names, "Mike")

	ages := make([]int, 0)
	ages = append(ages, 14)
	ages = append(ages, 34)
	ages = append(ages, 41)
	ages = append(ages, 19)
	ages = append(ages, 28)

	genders := make([]string, 0)
	genders = append(genders, "female")
	genders = append(genders, "male")
	genders = append(genders, "female")
	genders = append(genders, "male")
	genders = append(genders, "male")

	persons := make([]Person, 0)

	for i := 0; i < len(names); i++ {
		person := Person{name: names[i], age: ages[i], gender: genders[i]}
		persons = append(persons, person)
	}

	for _, person := range persons {
		person.PrintName()
	}
}
