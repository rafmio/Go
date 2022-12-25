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
		io.WriteString(writer, sh.message)
	}

func main() {
	err := http.ListenAndServe(":5000", StringHandler{ message: "Hello, World"})
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}

// type Handler interface { ServeHTTP(ResponseWriter, *Request)}
// A Handler responds to an HTTP request
// ServeHTTP()
// ServeHTTP(ResponseWriter, *Request)
// ServeHTTP should write reply headers and data to the ResponseWriter and then
// return
//
// ListenAndServe()
// func ListenAndServe(addr string, handler Handler) error
// listen on the TCP network address 'addr' and then calls Serve with handler
// requests on incoming connections
//
// This function starts listening for HTTP requests on a specified
// address and passes requests onto the specified handler.
//
// When a request arrives, it is passed onto a hanler, which is responsible
// for producing a response.
// Hanlder must implement the Hanlder interface
// (with method inside ServeHTTP(writer, request))
//
//
// func ListenAndServe(addr string, handler Handler) error {
// 	server := &Server{Addr: addr, Handler: handler}
// 	return server.ListenAndServe()
// }
//
// ResponseWriter
// type ResponseWriter interface {
// 		Header() Header
// 		Write([]byte)(int, err)
//  	WriteHeader(statusCode int)
// }
//
// Type Header
// type Header map[string][]string
// A Header represents the key-value pairs in an HTTP header
// The keys should be in canonical form
// Hanlder() returns the header map that will be sent by WriteHeader
// The Header map also is the mechanism with which Handlers cat set HTTP trailers
//
// Request
// type Request struct {
// 	Method string
// 	URL *url.URL
// 	...
// 	Host string
// 	...
// }
