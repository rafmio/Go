package main

import (
	"net/http"
	"io"
)

type StringHanlder struct{
	message string
}

func (sh StringHanlder) ServeHTTP(writer http.ResponseWriter,
request *http.Request) {
	Printfln("request.Method: %v", request.Method)
	Printfln("Request for %v", request.URL.Path)
	Printfln("request.Proto: %v", request.Proto)
	for key, val:= range request.Header {
		Printfln("[%v] : %v", key, val)
	}

	for key, val := range request.Trailer {
		Printfln("Trailer : [%v] : %v", key, val)
	}
	io.WriteString(writer, sh.message)
}

func main() {
	http.Handle("/message", StringHanlder{"Hello Huggy-Wuggy"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	err := http.ListenAndServe(":5000", nil)
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}

}

// HandlerFunc
