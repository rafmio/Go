package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

var tpl = template.Must(template.New("PersonTemplate").Parse(`{{ define "PersonTemplate"}}
	{{ if lt .Age 18 }}
		{{ .Name }} is over 18 years old.
	{{ else }}
		{{ .Name }} isn't over 18 yet.
	{{ end }}
{{ end }}`))

func main() {
	person := Person{Name: "John", Age: 19}
	err := tpl.Execute(os.Stdout, person)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

// eq arg1 arg2 — arg1 == arg2
// ne arg1 arg2 — arg1 != arg2
// lt arg1 arg2 — arg1 < arg2
// le arg1 arg2 — arg1 <= arg2
// gt arg1 arg2 — arg1 > arg2
// ge arg1 arg2 — arg1 >= arg2
