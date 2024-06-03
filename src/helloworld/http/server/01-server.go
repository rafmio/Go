// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"fmt"
	"net/http"
)

type HttpHandler struct{}

func (c HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("Hey there, tosy-bosy\n")

	res.Write(data)
	res.Write([]byte(req.Method))
	res.Write([]byte(req.Proto))
	fmt.Fprintf(res, "\n%s\n", "------")
}

func main() {
	handler := HttpHandler{} // an instance of HttpHandler struct
	http.ListenAndServe(":9000", handler)
}
