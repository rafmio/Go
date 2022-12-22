// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("")
	e := `Name,Scores
{{range $k, $v := .}}{{$k}} {{range $s := $v}},{{$s}}{{end}}
{{end}}`
	t.Parse(e)
	t.Execute(os.Stdout, map[string][]int{
		"Mike":  {88, 77, 99},
		"Betty": {54, 96, 78},
		"Jake":  {89, 67, 93},
	})
}
