package main

import (
  "html/template"
  // "os"
)

func main() {
  allTemplates, err := template.ParseGlob("templates/*.html")
  if (err == nil) {
    for _, t := range allTemplates.Templates() {
      Printfln("Template name: %v", t.Name())
    }
  } else {
    Printfln("Error: %v", err.Error())
  }
}
