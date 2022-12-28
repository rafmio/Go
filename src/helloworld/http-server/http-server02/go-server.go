// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"net/http"
	"io"
	"fmt"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "Hello")

	// write `World` using `fmt.Fprint` functin
	fmt.Fprint(res, " World")

	// write `❤️` using simple `Write` call
	res.Write([]byte(" ❤️"))
}


func main() {
	// create a new handler
	handler := HttpHandler{} // an instance of HttpHandler struct
	http.ListenAndServe(":9000", handler)
	// any incoming connection will trigger the ServeHTTP method
	// handler - это и есть наш ResponseWriter в дальнейшем?
}
