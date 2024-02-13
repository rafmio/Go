package main

import (
	"fmt"
	"os"
	"text/template"
)

type Inventory struct {
	Meterial string
	Count    int
}

func main() {
	sweaters := &Inventory{"wool", 17}
	forParse := "{{.Count}} items are made of {{.Meterial}}.\n"

	tmpl, err := template.New("test").Parse(forParse)
	if err != nil {
		fmt.Println("Error parsing:", err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		fmt.Println("Error os.Execute:", err)
	}
}
