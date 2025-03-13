package main

import (
  "html/template"
  "net/http"
  "strconv"
)

type Context struct {
  Request *http.Request
  Data []Product
}

var htmlTemplates *template.Template

func HandleTemplateRequest(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  if (path == "") {
    path = "products.html"
  }

  t := htmlTemplates.Lookup(path)
  if (t == nil) {
    http.NotFound(w, r)
  } else {
    err := t.Execute(w, Context {r, Products})
    if (err != nil) {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }
}

func init() {
  var err error
  htmlTemplates = template.New("all")
  htmlTemplates.Funcs(map[string]interface{} {
    "intVal": strconv.Atoi,
  })
  htmlTemplates, err = htmlTemplates.ParseGlob("templates/*.html")
  if (err == nil) {
    http.Handle("/templates/", http.StripPrefix("/templates/",
      http.HandlerFunc(HandleTemplateRequest)))
  } else {
    panic(err)
  }
}
