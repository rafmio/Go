// Creating text templates
package main

import (
	"os"
	"strings"
	"text/template"
)

func GetCategiries(products []Product) (categories []string) {
	catMap := map[string]string {}
	for _, p := range products {
		if (catMap[p.Category] == "") {
			catMap[p.Category] = p.Category
			categories = append(categories, p.Category)
		}
	}
	return
}

func Exec(t *template.Template) error {
	productMap := map[string]Product {}
	for _, p := range Products {
		productMap[p.Name] = p
	}
	return t.Execute(os.Stdout, &productMap)
}

func main() {
	allTemplates := template.New("allTemplates")
	allTemplates.Funcs(map[string]interface{} {
		"getCats": GetCategiries,
		"lower": strings.ToLower,
	})
	allTemplates, err := allTemplates.ParseGlob("templates/*.txt")

	if (err == nil) {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
