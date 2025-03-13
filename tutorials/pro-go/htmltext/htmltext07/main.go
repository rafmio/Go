package main

import (
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	allTemplated, err := template.ParseGlob("templates/*.html")
	if (err == nil) {
		selectedTemplated := allTemplated.Lookup("template.html")
		err = Exec(selectedTemplated)
	}
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}
