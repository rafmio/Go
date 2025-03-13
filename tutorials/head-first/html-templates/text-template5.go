package main

import (
    "os"
    "text/template"
    "log"
    "fmt"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func executeTemplate(text string, data interface{}) {
    tmpl, err := template.New("test_template").Parse(text)
    check(err)
    err = tmpl.Execute(os.Stdout, data)
    check(err)
}

func main() {
    var textOfTemplate string = "Dot is: {{.}}!\n"
    executeTemplate(textOfTemplate, "ABC")
    executeTemplate(textOfTemplate, 123.5)
    
    fmt.Println()
    templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
    executeTemplate(templateText, []string{"do", "re", "mi"})
    
    fmt.Println()
    templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
    executeTemplate(templateText, []float64{1.25, 0.99, 27})
    
    fmt.Println()
    templateText = "start {{if .}}Dot is true!{{end}} finish\n"
    executeTemplate(templateText, true)
    executeTemplate(templateText, false)
}
