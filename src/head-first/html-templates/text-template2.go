package main

import (
    "os"
    "log"
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
        Thank you for the lovely {{.}}.
        {{end}}
        Best wishes, 
        Josie
    `
    
    type Recipient struct {
        Name, Gift string
        Attended bool
    }
    
    var recipients = []Recipient{
        {"Aunt Mildred", "bone china tea set", true},
        {"Uncle John", "moleskin pants", false},
        {"Cousin Rodmey", "", false},
        {"Dyadya Valera", "fake vagina", true},
    }
    
    t := template.Must(template.New("letter").Parse(letter))
    
    for _, r := range recipients {
        err := t.Execute(os.Stdout, r)
        if err != nil {
            log.Println("executing template:", err)
        }
    }
}

// https://pkg.go.dev/text/template#Template
