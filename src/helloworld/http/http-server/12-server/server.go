package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server is listening on port 8181")
	http.ListenAndServe(":8181", http.FileServer(http.Dir("static")))
}
