package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  response, err := http.Get("https://example.com")
  if err != nil {
    log.Fatal(err)
  }

  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(body))
  fmt.Printf("Type of body: %T\n", body)
  fmt.Println(response.Status)
}

// http.Get возвращает объект http.Response
// http.Response - структура с полем Body (содержание страницы)
// Body поддерживает интерфейс ReadCloser уровня пакета
// ReadCloser поддерживает методы Read (для чтения)и Close (для освобождения
// сетевого подключения)

// Тело ответа передается функции ReadAll, который читает содержимое и
// возвращает в виде слайса (slice) значений byte
