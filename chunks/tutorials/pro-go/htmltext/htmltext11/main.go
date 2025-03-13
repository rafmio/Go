package main

import (
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	allTemplates, err := template.ParseGlob("templates/*.html")
	if (err == nil) {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
