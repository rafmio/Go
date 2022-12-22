// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"math"
	"os"
	"text/template"
)

func main() {
	t := template.New("")
	t.Parse(`Keeping just 2 decimals of π: {{printf "%.2f" .}}`)
	// t.Parse(`Keeping just 2 decimals of π: {{.}}`)
	t.Execute(os.Stdout, math.Pi)
}
