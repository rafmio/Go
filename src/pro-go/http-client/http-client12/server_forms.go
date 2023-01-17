package main

import (
  "net/http"
  "fmt"
  "io"
)

func init() {
  http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10_000_000)
    if (err == nil) {
      for name, vals := range r.MultipartForm.Value {
        fmt.Fprintf(w, "Field %v: %v\n", name, vals)
      }
      for name, files := range r.MultipartForm.File {
        for _, file := range files {
          fmt.Fprintf(w, "File %v: %v\n", name, file.Filename)
          if f, err := file.Open(); err == nil {
            defer f.Close()
            io.Copy(w, f)
          }
        }
      }
    } else {
      fmt.Fprintf(w, "Cannot parse form %v", err.Error())
    }
  })
}
