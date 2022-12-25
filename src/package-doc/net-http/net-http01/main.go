// https://habr.com/ru/company/skillbox/blog/446454/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Skillbox on habr.com!")
	})
	http.ListenAndServe(":5000", nil)
}
