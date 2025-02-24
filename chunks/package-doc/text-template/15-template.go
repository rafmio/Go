package main

import (
	"fmt"
	"os"
	"text/template"
)

type Student struct {
	ID   uint
	Name string
}

func main() {
	stu := Student{0, "Jason"}
	tmpl, err := template.New("test").Parse("The name for student {{ .ID }} is {{ .Name }}")
	if err != nil {
		fmt.Println("Error creating New template:", err.Error())
	} else {
		err = tmpl.Execute(os.Stdout, stu)
		fmt.Println()
	}

	tmplstr := `{{ define "stu_name" }}
	Student name is {{ .Name }}	
	{{ end }}
	{{ define "stu_info" }}
	{{ template "stu_name" . }}
	{{ end }}
	`

	tmpl2, err := template.New("test2").Parse(tmplstr)
	if err != nil {
		fmt.Println("Error creating New template", err.Error())
	} else {
		err = tmpl2.Execute(os.Stdout, stu)
	}
}
