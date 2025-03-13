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

// The templates must be loaded so that the file that contains the block action
// is loaded before the file that contains the define action that redefines the
// template.
// When the templates are loaded, the template defined in the list.html
// file redefines the template named body so that the content in the list.html
// file replaces the content in the template.html file.
