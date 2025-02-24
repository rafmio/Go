package main

import (
  "io"
  "log"
  "net/http"
)

func main() {
  h1 := func(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello form HandleFunc #1!\n")
  }

  h2 := func(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "Hello from a HandleFunc #2!\n")
  }

  http.HandleFunc("/", h1)
  http.HandleFunc("/endpoint", h2)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

// https://pkg.go.dev/net/http#HandleFunc

// type ResponseWriter interface {
//    Header() Header
//    Write([]byte)(int, error)
//    WriteHeader(statusCode int)
// }
// ResponseWriter используется HTTP-обработчиком для создания HTTP-ответа
//
// type Header map[string][]string
// -----------------------------------------------------------------------
// func HandleFunc(pattern string, handler func(ResponseWriter, *Requset))
// Функция HanldeFunc регистрирует обработчик функции для заданного
// шаблона DefaultServeMux

// func ListenAndServe(addr string, hanlder Handler) error
// Функция ListenAndServe прослушивает сетевой адрес TCP 'addr' а затем
// вызывает Serve с обработчиком для обработки запроса на взодящие соединения

// type Request struct {
//   Method string
//   URL *url.URL
//   Proto string
//   ProtoMajor int
//   ProtoMinor int
//   Header Header
//   Body io.ReadCloser
//   GetBody func() (io.ReadCloser, error)
//   ContentLength int64
//   TransferEncoding []string
//   Close bool
//   Host string
//   Form url.Values
//   PostForm url.Values
//   MultipartForm *multipart.Form
//   Trailer Header
//   RemoteAddr string
//   RequestURI string
//   TLS *tls.ConnectionState
//   Cancel <-chan struct{}
//   Response *Response
// }
// Request представляет собой HTTP-запрос, полученный сервером
// или подлежащий отправке клиентом
