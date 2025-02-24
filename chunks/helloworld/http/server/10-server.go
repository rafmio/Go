// https://medium.com/rungo/secure-https-servers-in-go-a783008b36da
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Hello-mello, Tosy-Bosy")
		fmt.Fprintln(res, "Host:", req.Host)
		fmt.Fprintln(res, "Proto:", req.Proto)
	})

	log.Fatal(http.ListenAndServeTLS(":9000", "home/raf/ssl-tls/certs/cert.pem", "home/raf/ssl-tls/certs/key.pem", nil))
}
