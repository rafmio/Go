// https://golangify.com/template-engine
package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
  "strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  homePagePath := "/home/raf-pc/Go/src/package-doc/html-template/home.page.tmpl"
  ts, err := template.ParseFiles(homePagePath)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
    return
  }

}
