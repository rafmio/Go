package main

import (
    "text/template"
    "os"
    "fmt"
)

type Joke struct {
    Who string
    Punchline string
}

func main() {
    t := template.New("Knock Knock Joke")
    text := `Knock Knock\nWho's there?
             {{.Who}}
             {{.Who}} who?
             {{.Punchline}}
            `
    t.Parse(text)

    jokes := []Joke{
        {"Etch", "Bless you!"},
        {"Cow goes", "No, cow goes moo!"},
    }

    for _, joke := range jokes {
        t.Execute(os.Stdout, joke)
    }
    fmt.Println()
}

// https://code.tutsplus.com/ru/tutorials/text-generation-with-go-templates--cms-30441
