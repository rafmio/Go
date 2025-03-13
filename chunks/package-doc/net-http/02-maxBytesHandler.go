package main

import (
	"fmt"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error loading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "File uploaded successfully")
}

func main() {
	http.Handle(
		"/upload",
		http.MaxBytesHandler(
			http.HandlerFunc(uploadHandler),
			5*1024*1024,
		),
	)

	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

// curl -X POST -F "file=@/path/to/file" http://localhost:8080/upload
