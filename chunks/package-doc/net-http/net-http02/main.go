// https://dev-gang.ru/article/go-velikolepnyi-modul-nethttp-u69osny6tr/
package main

import (
  "fmt"
  "net/http"
)

const portNumber = ":8080"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("We're alive!"))
  fmt.Fprintf(w, "\nAnd we are healthy")
}

func main() {
  http.HandleFunc("/", HomePageHandler)
  fmt.Printf("Starting application on port %v\n", portNumber)
  http.ListenAndServe(portNumber, nil)
}

// Метод HandleFunc присоединяет путь к обработчику и каждая функция обработчика
// должна принять HTTP request (указатель на него), и writer, который отвечает
// за отправку обратно ответа через сеть клиенту
