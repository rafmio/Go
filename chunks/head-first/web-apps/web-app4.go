package main

import (
  "log"
  "net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
  message := []byte("Hello-mello, web-meb!")
  _, err := writer.Write(message)
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/hello", viewHandler)
  err := http.ListenAndServe("localhost:8080", nil)
  log.Fatal(err)
}

// ListenAndServe запускает веб-сервер с прослушиванием запросов 
