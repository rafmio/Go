package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := &Inventory{"wool", 17}
	forParse := "{{.Count}} items are made of {{.Material}}\n"

	tmpl, err := template.New("test").Parse(forParse)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
