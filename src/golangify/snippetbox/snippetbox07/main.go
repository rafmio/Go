package main

import (
  "log"
  "net/http"
  "fmt"
  "strconv"

)

// home ----------------------------------------------------------------------
func home(w http.ResponseWriter, r *http.Request) {
  if (r.URL.Path != "/") {
    http.NotFound(w, r)
    return
  }

  w.Write([]byte("HOME\n"))
  w.Write([]byte("Hello inside of SnippetBox\n"))
}

// show snippet --------------------------------------------------------------
func showSnippet(w http.ResponseWriter, r *http.Request) {
  // Извлекаем значение параметра id из URL и попытаемся конвертировать строку
  // в  integer, испльзуя strconv.Atoi()

  log.Println("id: ", r.URL.Query().Get("id"))
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if (err != nil || id < 1) {
    http.NotFound(w, r)
    return
  }

  fmt.Fprint(w, "Displaying selected snippet of ID %d ...", id)
}


// new snippet ---------------------------------------------------------------
// Handler for creating the new Snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
  // Используем r.Method для проверки использует ли запрос метод POST
  if (r.Method != http.MethodPost) {
    // Используем метод Header().Set() для добавления заголовка 'Allow: POST'
    // в карту HTTP-заголовков. Ответ будет включать [Allow]: POST
    w.Header().Set("Allow", http.MethodPost)

    // w.WriteHeader(http.StatusMethodNotAllowed) // 405
    // w.Write([]byte("Method is not allowed\n")) // Method запрещен
    // Вместо закомментированного кода используем http.Error():
    http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed) // 405
    // Все осталось прежним, кроме того, что теперь вызов http.ResponseWriter
    // мы передает под ответственность http.Error(), которая за нас отправляет
    // ответ пользователю
    return
  }
  w.Write([]byte("Creating a new snippet\n"))
}

// main ----------------------------------------------------------------------
func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  log.Println("Launching the web-server on http://127.0.0.1:4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}
