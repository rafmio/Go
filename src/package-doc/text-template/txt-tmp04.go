// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

type Family struct {
	Father        Person
	Mother        Person
	ChildrenCount int
}

type Person struct {
	Name   string
	Gender string
}

func main() {
	family := Family{
		Father:        Person{"Tarzan", "male"},
		Mother:        Person{"Jane", "female"},
		ChildrenCount: 2,
	}

	t := template.New("Father")
	text := "The father's name is {{.Father.Name}}\n"
	t.Parse(text)
	t.Execute(os.Stdout, family)
}
