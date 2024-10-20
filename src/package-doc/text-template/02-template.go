// https://pkg.go.dev/text/template#example-Template
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}!
I will use {{.}} wisely.
{{end}}
Best wishes,
Josie
`

	type Recipient struct {
		Name, Gift string
		Attended   bool
	}

	var recipient = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "shirt", true},
	}

	t := template.Must(template.New("letter").Parse(letter))

	for _, r := range recipient {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("execute error: ", err)
		}
	}

	err := t.Execute(os.Stdout, recipient[0])
	if err != nil {
		log.Println("execute error: ", err)
	}

	fmt.Println("template's name is: ", t.Name())
}
