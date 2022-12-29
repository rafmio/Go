// https://medium.com/rungo/secure-https-servers-in-go-a783008b36da
// package main
//
// import (
// 	"net/http"
// 	"fmt"
// 	"log"
// )
//
// func main() {
// 	// handle `/` route
// 	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
// 		fmt.Fprint(res, "Hello World")
// 	})
//
// 	// run server on port 9000
// 	log.Fatal(http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil))
// }

package main

import (
	"net/http"
	"fmt"
	"log"
	"crypto/tls"
)

func main() {
	// generate a `Certificate` struct
	cert, err := tls.LoadX509KeyPair("/home/raf/ssl-tls/certs/cert.pem", "/home/raf/ssl-tls/certs/key.pem")
	if (err != nil) {
		fmt.Println(err.Error())
	}
	s := &http.Server{
		Addr: ":9000",
		Handler: nil,			// user `http.DefaultServeMux`
		TLSConfig: &tls.Config {
			Certificates: []tls.Certificate{cert},
		},
	}

	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request){
		fmt.Fprint(res, "Hello, Custom World!")
	})

	// run server on port 9000
	log.Fatal(s.ListenAndServeTLS("", ""))
}
