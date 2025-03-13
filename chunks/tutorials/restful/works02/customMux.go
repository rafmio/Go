package main

import (
  "fmt"
  "math/rand"
  "net/http"
)

func main() {
  newMux := http.NewServeMux()
  newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, rand.Float64())
  })
  newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, rand.Intn(100))
  })

  fmt.Println("Running listening for http... ")

  err := http.ListenAndServe(":8000", newMux)
  if err != nil {
    fmt.Println("runnig listening server: ", err.Error())
  }
}
