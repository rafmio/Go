package main

import (
  "fmt"
  "net/http"
  "strconv"
  "html/template"
  "log"
)
// --------------------------------------------------------------------------
func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  files := []string{
    "./ui/html/home.page.tmpl",
    "./ui/html/base.layout.tmpl",
    "./ui/html/footer.partial.tmpl",
  }

  ts, err := template.ParseFiles(files...)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", http.StatusInternalServerError) // 500
    return
  }

  err = ts.Execute(w, nil)
  if err != nil  {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)  // 500
  }

  // w.Write([]byte("Hello inside of SnippetBox"))
}

// ---------------------------------------------------------------------------

func showSnippet(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(w, r)
  }

  fmt.Fprint(w, "Displaying selected snippet with ID %d...", id)
}

// ---------------------------------------------------------------------------

func createSnippet(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost) // К перечню заголовков добавили "Allow"
    http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
    return
  }
  w.Write([]byte("Creating a new snippet"))
}
