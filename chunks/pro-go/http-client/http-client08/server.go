package main

import (
  "encoding/json"
  "fmt"
  "io"
  "net/http"
  "os"
)

func init() {
  http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./index.html")
  })

  http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Products)
  })

  http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "Method: %v\n", r.Method)
    for header, vals := range r.Header {
      fmt.Fprintf(w, "Header: %v: %v\n", header, vals)
    }
    fmt.Fprintln(w, "----------")
    data, err := io.ReadAll(r.Body)
    if (err == nil) {
      if len(data) == 0 {
        fmt.Fprintln(w, "No body")
      } else {
        w.Write(data)
      }
    } else {
      fmt.Fprintf(os.Stdout, "Error reading body: %v\n", err.Error())
    }
  })
}
