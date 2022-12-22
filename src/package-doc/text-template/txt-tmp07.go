// Pipelines
// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("")
	t.Parse(`{{ call . | len | printf "%d items" }}`)
	t.Execute(os.Stdout, func() string { return "abc" })
}

// First, the call function executes the function pass to Execute().
// Then the len function returns the length of the result of the input function,
// which is 3 in this case.
// Finally, the printf function prints the number of items.
