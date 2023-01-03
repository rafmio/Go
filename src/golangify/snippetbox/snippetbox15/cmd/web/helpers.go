package main

import (
  "fmt"
  "net/http"
  "runtime/debug"
)

// serverError записывает сообщение об ошибке в errorLog и затем
// отправляет пользователю ошибку 500 "Внутненняя ошибка сервера"
func (app *application) serverError(w http.ResponseWriter, err error) {
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
  app.errorLog.Println(trace)

  http.Error(w, http.StatusText(http.StatusInernalServerError), http.StatusInernalServerError)
}

// clientError отправляет определенный код состояния + описание пользователю
func (app *application) clientError(w http.ResponseWriter, status int) {
  http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
  app.clientError(w, http.StatusNonFound)
}
