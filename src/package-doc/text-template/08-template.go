// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("08-template.go")
	text := `
{{ $F := .FirstName | printf "%q"}}
{{ $L := .LastName  | printf "%q"}}
Normal: {{$F}} {{$L}}
Reverse: {{$L}} {{$F}}
`
	t.Parse(text)

	type Person struct {
		FirstName string
		LastName  string
	}

	person := &Person{
		FirstName: "Gigi",
		LastName:  "Sayfan",
	}

	t.Execute(os.Stdout, person)
}
