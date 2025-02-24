// https://medium.com/rungo/creating-a-simple-hello-world-http-server-in-go-31c7fd70466e
package main

import (
	"fmt"
	"io"
	"net/http"
)

type HttpHandler struct{}

func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello-mello")

	fmt.Fprint(res, " World")

	res.Write([]byte(" ❤️\n"))
}

func main() {
	handlder := HttpHandler{}
	http.ListenAndServe(":9000", handlder)
}

// curl -X GET http://localhost:9000/
