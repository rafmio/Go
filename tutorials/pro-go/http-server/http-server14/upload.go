package main

import (
  "net/http"
  "io"
  "fmt"
)

func HandleMultipartForm(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Name: %v, City: %v\n", r.FormValue("name"),
  r.FormValue("city"))
  fmt.Fprintln(w, "-----------")

  file, header, err := r.FormFile("files")
  if (err == nil) {
    defer file.Close()
    fmt.Fprintf(w, "Name: %v, Size: %v\n", header.Filename, header.Size)
    for k, v := range header.Header {
      fmt.Fprintf(w, "Key: %v, Value: %v\n", k, v)
    }
    fmt.Fprintln(w, "-----------")
    io.Copy(w, file)
  } else {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func init() {
  http.HandleFunc("/forms/upload", HandleMultipartForm)
}
