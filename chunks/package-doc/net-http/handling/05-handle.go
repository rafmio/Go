package main

import (
	"encoding/json"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Hello, World!",
		"status":  "success",
	})

	// fmt.Fprintf(w, "Hello-mello from MyHandler!")
}

func main() {
	handler := MyHandler{}
	http.Handle("/custom", &handler)

	http.ListenAndServe(":8080", nil)
}
