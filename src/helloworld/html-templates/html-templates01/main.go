// https://habr.com/ru/post/475390/
package main

import (
  "net/http"
  "html/template"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
  tpl.Execute(w, nil)
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", indexHandler)

  http.ListenAndServe(":4000", mux)
}
