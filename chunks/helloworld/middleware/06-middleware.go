// Example provided by DeepSeek
package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello-mello from loggingMiddleware()\n"))
		log.Println("The request has been received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello-mello from 'helloHandler()\n"))
}

func main() {
	/*
		 	My notes:
			The second argument of the http.Handle function is the http.Handler data type.
			The loggingMiddleware function takes http.Handler as an argument and returns it.
			With the http.HandlerFunc function, we convert the helloHandler() function to
			the http data type.Hanler
	*/
	http.Handle("/", loggingMiddleware(http.HandlerFunc(helloHandler)))
	log.Println("The server run on :8080")
	http.ListenAndServe(":8080", nil)
}
