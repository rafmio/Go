package main

import (
    "text/template"
    "os"
    "log"
    "fmt"
)

func check(err error){
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    templateText := "Template start\nAction: {{.}}\nTemplate end\n\n"
    tmpl, err := template.New("Hohoho").Parse(templateText)
    check(err)
    err = tmpl.Execute(os.Stdout, "ABC")
    check(err)
    err = tmpl.Execute(os.Stdout, 42)
    check(err)
    err = tmpl.Execute(os.Stdout, true)
    check(err)
    
    fmt.Println(tmpl)
    fmt.Printf("tmpl: %#v\n", tmpl)
    fmt.Println()
    
    tmplcpy := tmpl 
    err = tmpl.Execute(os.Stdout, tmplcpy)
    check(err)
}
