package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.ParseFiles("./html/index.html")
  if err != nil {
    fmt.Printf("Error parsing files: %v\n", err.Error())
  }
  tmpl.Execute(w, Entries)

}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handler)

  for i, v := range Entries {
    fmt.Println(i, v.LoadingDate, v.Truck)
  }

  http.ListenAndServe(":5000", mux)
}
