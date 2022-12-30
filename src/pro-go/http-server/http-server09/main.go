package main

import (
	"net/http"
	"io"
	"strings"
	"fmt"
)

type StringHandler struct {
	message string
}

type HomePage struct {
	pattern	 string
	host string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
}

func (hp HomePage) ServeHTTP(writer http.ResponseWriter,
request *http.Request) {
	Printfln("Handler for HomePage is in action")
	io.WriteString(writer, "HomePage handler is in action\n")
	io.WriteString(writer, hp.pattern)
	io.WriteString(writer, "\n")
	io.WriteString(writer, hp.host)
	fmt.Fprint(writer, "\n\nUsing fmt.Fprint()")
}

func HTTPSRedirect(writer http.ResponseWriter, request *http.Request) {
	Printfln("request.Host: '%v'", request.Host)
	Printfln("request.URL.Path: '%v'", request.URL.Path)

	host := strings.Split(request.Host, ":")[0]
	Printfln("host: '%v'", host)

	target := "https://" + host + ":5500" + request.URL.Path
	Printfln("target: '%v'", target)

	if len(request.URL.RawQuery) > 0 {
		target += "?" + request.URL.RawQuery
	}
	Printfln("request.URL.RawQuery: %v", request.URL.RawQuery)

	Printfln("target: '%v'", target)

	http.Redirect(writer, request, target, http.StatusTemporaryRedirect) // Кому переадресует?
}

func main() {
	http.Handle("/message", StringHandler{message: "Hello, redirected World"})
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	http.Handle("/homepage", HomePage{pattern: "homepage", host: "localhost"})

	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.crt", "certificate.key", nil)
		if (err != nil) {
			Printfln("HTTPS Error: %v", err.Error())
		}
		}()

		err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
		if (err != nil) {
			Printfln("Error: %v", err.Error())
		}
}
