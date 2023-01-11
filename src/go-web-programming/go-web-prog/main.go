package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, GWP, World! %s", r.URL.Path[1:])
  fmt.Fprintf(w, "\nr.URL.Path[0]: %v", r.URL.Path[0])
  
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
