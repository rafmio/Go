package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello-mello from the 'helloWorld()' func")
}

func main() {
	http.Handle("/hello", http.HandlerFunc(helloWorld))

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}
