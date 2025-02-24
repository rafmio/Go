// https://dev-gang.ru/article/go-velikolepnyi-modul-nethttp-u69osny6tr/
package main

import (
  "fmt"
  "net/http"
)

const portNumber = ":8080"

func HomePageHandler(writer http.ResponseWriter, r *http.Request)  {
  fmt.Fprintf(writer, "Hello %v\n", r.UserAgent())
  fmt.Fprintf(writer, "You are at %v\n", r.URL.Path)
  fmt.Fprintf(writer, "You have sent some query params to: %v\n", r.URL.Query())
  fmt.Fprintf(writer, "The method you used is: %v", r.Method)
}

func main() {
  http.HandleFunc("/", HomePageHandler)
  fmt.Printf("Starting application on port %v\n", portNumber)
  http.ListenAndServe(portNumber, nil)
}
