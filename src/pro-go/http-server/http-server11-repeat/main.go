package main

import (
  "net/http"
  "io"
)

type StringHandler struct {
  message string
}

func (sh StringHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  Printfln("Request for %v", r.URL.Path)
  switch r.URL.Path {
  case "/favico.ico":
    http.NotFound(w, r)
  case "/message":
    io.WriteString(w, sh.message)
  default:
    http.Redirect(w, r, "/message", http.StatusTemporaryRedirect)
  }
}

func main() {
  err := http.ListenAndServe(":5000", StringHandler {message: "<p>Susi-Musi, Tosi-Bosi</p>"})
  if (err != nil) {
    Printfln("Error: %v", err.Error())
  }
}
