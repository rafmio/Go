// Pipelines
// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("07-template.go")
	text := `{{ call . | len | printf "%d items"}}`
	t.Parse(text)

	t.Execute(os.Stdout, func() string { return "abcd" })
	t.Execute(os.Stdout, func() string { return "Kissy-Missy" })

	tt := template.New("Second template")
	text2 := `{{ call . | printf "%s"}}`
	tt.Parse(text2)
	tt.Execute(os.Stdout, func() string { return "Huggy-Wuggy" })
}

// First, the call function executes the function pass to Execute().
// Then the len function returns the length of the result of the input function,
// which is 3 in this case.
// Finally, the printf function prints the number of items.
