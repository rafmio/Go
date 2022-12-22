// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("")
	t.Parse(`
        {{ $F := .FirstName | printf "%q"}}
        {{ $L := .LastName  | printf "%q"}}
        Normal:  {{$F}} {{$L}}
        Reverse: {{$L}} {{$F}}
    `)

	t.Execute(os.Stdout, struct {
		FirstName string
		LastName  string
	}{
		"Gigi",
		"Sayfan",
	})
}
