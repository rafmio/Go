// https://metanit.com/go/web/2.1.php
package main

import (
  "fmt"
  "net/http"
  "html/template"
)

type ViewData struct {
  Title string
  Message string
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    data := ViewData {
      Title: "World Cup",
      Message: "FIFA will never regret it",
    }
    tmpl, _ := template.ParseFiles("./templates/index.html")
    tmpl.Execute(w, data)
  })

  fmt.Println("Serverer is listening...")
  http.ListenAndServe(":8181", nil)
}
