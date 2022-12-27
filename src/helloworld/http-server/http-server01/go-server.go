// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"net/http"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// create response binary data
	data := []byte("Hello from medium.com, tosy-bosy\n") // slice of bytes

	// write `data` to response
	res.Write(data)
	res.Write([]byte(req.Method))
	res.Write([]byte(req.Proto))
}

func main() {
	// create a new handler
	handler := HttpHandler{} // an instance of HttpHandler struct
	http.ListenAndServe(":9000", handler)
}
