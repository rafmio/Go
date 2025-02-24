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

// Handler for creating the new Snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
  // Используем r.Method для проверки использует ли запрос метод POST
  // или нет.
  // http.MethodPost является строкой и содержит текст "POST"
  if (r.Method != http.MethodPost) {
    // Используем метод Header().Set() для добавления заголовка 'Allow: POST'
    // в карту HTTP-заголовков. Первый параметр - название заголовка,
    // второй - значение заголовка
    w.Header().Set("Allow", http.MethodPost)

    // WARNING: сначала вызываем w.Header().Set(), и уже потом остальные методы,
    // (w.WriteHeader или w.Write), иначе изменения не будут учтены

    w.WriteHeader(http.StatusMethodNotAllowed) // 405
    w.Write([]byte("GET-Method is not allowed\n")) // GET-Method запрещен
    return
  }
  w.Write([]byte("Creating a new snippet\n"))
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
