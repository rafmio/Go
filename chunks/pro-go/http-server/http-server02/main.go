package main

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

// ServeHTTP записывает заголовки ответа и данные в ResponseWriter и возвращает
func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	Printfln("Method: %v", request.Method)
	Printfln("URL: %v", request.URL)
	Printfln("HTTP Version: %v", request.Proto)
	Printfln("Host: %v", request.Host)

	index := 0
	for name, val := range request.Header {
		Printfln("%v - Header: %v, Value: %v", index, name, val)
		index++
	}
	Printfln("%s", "-----------------------------------")
	io.WriteString(writer, sh.message)
}

func main() {
	err := http.ListenAndServe(":5000", StringHandler{message: "Hello-Mello"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
