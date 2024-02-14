// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl := template.New("10-template.go")
	text := `Name    Scores
{{range $k, $v := .}}{{$k}} {{range $s := $v}} {{$s}}{{end}}
{{end}}`
	tmpl.Parse(text)
	tmpl.Execute(os.Stdout, map[string][]int{
		"Mike":  {88, 77, 99},
		"Betty": {54, 96, 78},
		"Jake":  {89, 67, 93},
	})

	tmpl2 := template.New("11")
	text2 := `{{range $k, $v := .}}{{$k}} {{range $s := $v}}`
	tmpl2.Parse(text2)
	tmpl2.Execute(os.Stdout, map[string][]int{
		"Mike":  {88, 77, 99},
		"Betty": {54, 96, 78},
		"Jake":  {89, 67, 93},
	})
}
