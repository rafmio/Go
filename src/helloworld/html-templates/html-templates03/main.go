package main

import (
  "fmt"
  "net/http"
  "html/template"
)

type ViewData struct {
  Title string
  Users []string
}

func main() {
  data := ViewData {
    Title : "Users List",
    Users : []string {"Tom", "Bob", "Sam", "Valera", "Crocodilische"},
  }
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, data)
  })

  fmt.Println("Server is listening...")
  http.ListenAndServe(":8181", nil)
}
