package main

import (
  "log"
  "net/http"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс,
// текст "Hello inside of SnippetBox" как тело ответа
func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello inside of SnippetBox"))
}

func main() {
  // Используется ф-я http.NewServeMux() для инициализации нового роутера,
  // затем функцию home регистрируется как обработчик для URL-шаблона
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)

  // http.ListenAndServe - для запуска нового веб-сервера
  // Параметры:
  // 1.TCP-адрес для прослушивания
  // 2.Роутер
  log.Println("Launching web server on http://127.0.0.1:4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}

// ResponseWriter - методы для объединения HTTP ответа и возвращение его
// пользователю
// Request - содержит инфу о текущем запросе

// Маршутизатор рассматривает "/" как "catch-all"
