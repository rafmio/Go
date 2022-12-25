package main

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

// ServeHTTP записывает заголовки ответа и данные в ResponseWriter
func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	Printfln("Request for %v", request.URL.Path)

	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/message":
		io.WriteString(writer, sh.message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
}

func main() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello-Mello, Tosy-Bosy"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
