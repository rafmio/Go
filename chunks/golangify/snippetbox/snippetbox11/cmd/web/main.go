package main

import (
  "flag"
  "log"
  "net/http"
)

func main() {
  addr := flag.String("addr", ":4000", "HTTP net address ")
  flag.Parse()

  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet", showSnippet)
  mux.HandleFunc("/snippet/create", createSnippet)

  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/", http.StripPrefix("/static", fileServer))

  log.Printf("Launching server on %s", *addr)
  err := http.ListenAndServe(*addr, mux)
  log.Fatal(err)
}

// go run ./cmd/web -addr="127.0.0.1:9999"
// Значения по умолчанию
// Если запустить веб-приложение без флага -addr, то сервер будет использовать
// адрес :4000 - значение по умолчанию

// Справка:
// go run ./cmd/web -help
// Использование флага -help - ознакомление со списком доступных флагов
