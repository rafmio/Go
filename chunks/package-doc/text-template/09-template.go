// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("09-template.go")
	t.Parse(`{{ if . - }} {{ . }} {{ else }} No data is available {{ end }}`)
	t.Execute(os.Stdout, "42")
	t.Execute(os.Stdout, "")
}
