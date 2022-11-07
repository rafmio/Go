package main

import (
  "log"
  "net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
  message := []byte("Hello, web!")
  _, err := writer.Write(message)
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/hello", viewHandler)
  err := http.ListenAndServe("localhost:8080", nil) // runs web-server (local)
  log.Fatal(err)
}


// Сначала вызываем HandleFunc, только потом ListenAndServe
// HandleFunc для генерирования ответа вызывает viewHandler

// http.ResponseWriter используется для записи данных в ответ браузера
// http.Request - запрос браузер
// Значение Request не используется, но функция-обработчик все равно должна
// плучить его

// аргумент writer - значение для изменения ответа, отправляемого браузеру
// аргумент request - значение, предоставляющее запрос от браузера
