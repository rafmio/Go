package main

import (
  "fmt"
  "net/http"
  "runtime/debug"
)

// Записывает сообщение об ошибке в errorLog и отправляет пользователю ответ 500
func (app *application) serverError(w http.ResponseWriter, err error){
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
  app.errorLog.Output(2, trace)

  http.Error(w, http.StatusText(http.StatusInternalServerError),
  http.StatusInternalServerError)
}

// Отправляет определенный код состояние+описание пользователю
func (app *application) clientError(w http.ResponseWriter, status int) {
  http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
  app.clientError(w, http.StatusNotFound)
}
