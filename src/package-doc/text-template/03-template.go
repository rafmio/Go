// https://code.tutsplus.com/tutorials/text-generation-with-go-templates--cms-30441
package main

import (
	"log"
	"os"
	"text/template"
)

type Joke struct {
	Who       string
	Punchline string
}

func main() {
	t := template.New("Knock-knock yopta Joke!")
	text := `
Knock-knock, yopta!
Who's there?
{{.Who}}
{{.Who}} who?
{{.Punchline}}
`

	t.Parse(text)
	jokes := []Joke{
		{"Etch", "Bless-mess you!"},
		{"Cow goes", "No, cow goes moo!"},
		{"It's me", "Idi nahui!"},
	}

	for _, joke := range jokes {
		err := t.Execute(os.Stdout, joke)
		if err != nil {
			log.Fatal(err)
		}
	}
}
