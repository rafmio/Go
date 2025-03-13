package main

import (
	"fmt"
	"log"
	"net/http"
)

// Структура для хранения настроек конфигурации
type config struct {
	MaxBodyBytes int64 // Максимальный размер тела запроса в байтах
}

// Основная функция-обработчик
func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Промежуточный слой для логирования
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Промежуточный слой для проверки аутентификации
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем заголовок Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Функция для объединения middleware
func chainMiddleware(h http.Handler, m ...func(http.Handler) http.Handler) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

// Главная точка входа
func main() {
	// Настройки конфигурации
	cfg := &config{
		MaxBodyBytes: 1024 * 1024, // 1 MB
	}

	// Создаем основной обработчик
	mainHandler := http.HandlerFunc(handleRequest)

	// Применяем ограничение на размер тела запроса
	mainHandler = http.HandlerFunc(http.MaxBytesHandler(mainHandler, cfg.MaxBodyBytes).ServeHTTP)

	// Применяем middleware
	wrappedHandler := chainMiddleware(mainHandler, loggingMiddleware, authMiddleware)

	// Запускаем сервер
	log.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", wrappedHandler)
}

// $ curl -H "Authorization: Bearer valid-token" http://localhost:8080/

//  $ curl -X  GET http://localhost:8080/
