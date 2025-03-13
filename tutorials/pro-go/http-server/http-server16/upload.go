package main

import (
  "net/http"
  "io"
  "fmt"
)

func HandleMultipartForm(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(10_000_000)
  fmt.Fprintf(w, "Name: %v, City: %v\n",
  r.MultipartForm.Value["name"][0],
  r.MultipartForm.Value["city"][0])

  fmt.Fprintln(w, "--------------")

  for _, header := range r.MultipartForm.File["files"] {
    fmt.Fprintf(w, "Name: %v, Size: %v\n", header.Filename, header.Size)
    file, err := header.Open()
    if (err == nil) {
      defer file.Close()
      fmt.Fprintln(w, "--------------")
      io.Copy(w, file)
    } else {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  }
}

func init() {
  http.HandleFunc("/forms/upload", HandleMultipartForm)
}
