// https://www.golinuxcloud.com/golang-http/
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

// $ ss -ntlp | grep 8080
// $ curl -X GET localhost:8080
// $ curl -X GET localhost:8080/api
