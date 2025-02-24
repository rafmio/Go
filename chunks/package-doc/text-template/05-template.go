// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t := template.New("NewTemplate")
	t.Parse("Anything goes: {{.}}\n")
	t.Execute(os.Stdout, 1)
	t.Execute(os.Stdout, "two")
	t.Execute(os.Stdout, 3.0)
	t.Execute(os.Stdout, map[string]int{"four": 4})
	fmt.Println(t.Name())
}
