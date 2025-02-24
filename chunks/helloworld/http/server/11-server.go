// https://www.golinuxcloud.com/golang-web-server/
package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	for key, value := range r.Header {
		fmt.Println(key, value)
	}
}

func methodHandler(w http.ResponseWriter, r *http.Re)
