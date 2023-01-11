package main

import (
	"io"
	"net/http"
	"strings"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	io.WriteString(writer, sh.message)
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

	http.Redirect(writer, request, target, http.StatusTemporaryRedirect)
	// Запрос идет на адрес, порт и путь target с соотв.кодом
	// А там уже нужный порт слушает соответствующий (в горутине слушальщик, кот
	// запускается ListenAndServeTLS)
}

func main() {
	http.Handle("/message", StringHandler{"Hello-Mello, Tosy-Bosy StringHandler field"})
	http.Handle("/favico.ico", http.NotFoundHandler())
	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))

	fsHandler := http.FileServer(http.Dir("./static"))
	// Можно указывать путь к файлам и напрямую, но так можно легко позволить
	// запросам обращаться к файлам за пределы целевой папки, поэтому безопаснее
	// использовать Dir
	
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))

	go func() {
		err := http.ListenAndServeTLS(":5500", "certificate.crt", "certificate.key", nil)
		if err != nil {
			Printfln("HTTPS Error: %v", err.Error())
		}
	}()

	err := http.ListenAndServe(":5000", http.HandlerFunc(HTTPSRedirect))
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
