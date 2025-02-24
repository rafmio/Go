package server

import (
	"log"
	"net/http"

	"geniusapi/handlers"
)

func RunServer() {
	mux := http.NewServeMux()

	// Регистрация маршрута
	// mux.HandleFunc("/info", handlers.GetInfoHandler).Methods("GET")
	mux.HandleFunc("/info", handlers.GetInfoHandler)

	// Запуск сервера
	port := ":8000"
	log.Printf("the server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
