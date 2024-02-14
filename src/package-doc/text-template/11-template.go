package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	type Person struct {
		Name   string
		Scores []int
	}

	people := map[string][]int{
		"Mike":  {88, 77, 99},
		"Betty": {54, 96, 78},
		"Jake":  {89, 67, 93},
	}

	tmpl, err := template.New("people").Parse(`{{range $k, $v := .}}{{$k}} {{range $s := $v}} {{$s}}, {{end}}
	{{end -}}`)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(os.Stdout, people)
	if err != nil {
		log.Fatal(err)
	}
}
