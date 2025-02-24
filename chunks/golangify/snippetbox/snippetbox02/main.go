package main

import (
  "log"
  "net/http"
)

// Main page handler
func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello from inside of SnippetBox"))
}

// Handler for displaying the contents of the note
func showSnippet(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Displaying the note..."))
}

// Handler for creating a new note
func createSnippet(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Form for creating a new note..."))
}

func main() {
  // Регистирует два новых обработчика и соотв. URL-шаблоны в
  // маршутизаторе servemux
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  log.Println("Launching the web-server on http://127.0.0.1:4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
