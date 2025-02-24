package main

import (
  "log"
  "os"
  "text/template"
  "fmt"
)

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  text := "Here's my template\n"
  tmpl, err := template.New("test").Parse(text)
  check(err)
  err = tmpl.Execute(os.Stdout, nil)
  check(err)
  fmt.Printf("Type of tmpl: %T\n", tmpl)
  fmt.Println(text)
  fmt.Println(tmpl)
}
