package main

import (
  "log"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  if (r.URL.Path != "/") {
    http.NotFound(w, r)
    return
  }

  w.Write([]byte("Hello inside of SnippetBox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Displaying the note..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Form for creating the new note..."))
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  log.Println("Launching the web-server on http://127.0.0.1:4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
