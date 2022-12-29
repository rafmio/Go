package main

import (
	"net/http"
	"io"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func main() {
	http.Handle("/message", StringHandler{"Hello-mello, Kissy-Missy"})
	http.Handle("/favion.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.crt", "certificate.key", nil)
		if (err != nil) {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()

	err := http.ListenAndServe(":5000", nil)
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}
