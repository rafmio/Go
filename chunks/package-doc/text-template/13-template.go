// https://habr.com/ru/articles/792802/
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	type Conditions struct {
		Night     bool
		TimeOfDay string
	}

	conditions := Conditions{Night: true, TimeOfDay: "Night"}

	tmpl := template.New("tmpl1")
	text1 := `{{ .TimeOfDay }} is beautiful`
	tmpl.Parse(text1)

	err := tmpl.Execute(os.Stdout, conditions)
	if err != nil {
		fmt.Println("ERROR executing:", err)
	}
}
