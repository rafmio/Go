package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Обработчик GET-запросов
func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "I see you")
}

// Обработчик PUT-запроса /shutdown
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut || r.URL.Path != "/shutdown" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Закрываем сервер, останавливая процесс
	os.Exit(0)
}

func main() {
	http.HandleFunc("/", handleGet)
	http.HandleFunc("/shutdown", shutdownHandler)

	log.Println("Server is listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil && err.Error() != "http: Server closed" {
		log.Fatalf("Error starting the server: %v\n", err)
	}
}
