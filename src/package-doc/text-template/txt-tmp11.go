// https://networkbit.ch/golang-text-template/
package main

import (
	"os"
	"text/template"
)

type Vlans struct {
	Comment string
	Ids     []Ids
}

type Ids struct {
	Id   int
	Name string
}

func main() {
	config := `
{{- $.Comment -}}
{{range .Ids}}
vlan {{ .Id }} name {{ .Name -}}
{{end}}
`
	vlans := Vlans{
		Comment: "Vlans:",
		Ids: []Ids{
			{1, "One"},
			{2, "Two"},
			{3, "Three"},
		},
	}

	err := RunTemplate(config, vlans)
	if err != nil {
		panic(err)
	}
}

func RunTemplate(c string, v Vlans) error {
	tmpl, err := template.New("vlans").Parse(c)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, v)
	if err != nil {
		return err
	}

	return nil
}
