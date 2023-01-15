// https://metanit.com/go/web/1.3.php
package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/about.html")
  })
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/index.html")
  })
  fmt.Println("Server is listening...")
  http.ListenAndServe(":8181", nil)
}

// Также для отправки файлов можно использовать функцию http.ServeFile().
// Она отправляет единичный файл по определенному пути.
