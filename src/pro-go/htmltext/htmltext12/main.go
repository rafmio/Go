package main

import (
	"html/template"
	"os"
)

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	// To use this feature, the template files must be loaded in such order:
	allTempltes, err := template.ParseFiles("templates/template.html",
	"templates/list.html")
	if (err == nil) {
		selectedTemplated := allTempltes.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	} else {
		Printfln("Error: %v %v", err.Error())
	}
}
