package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}))
	http.ListenAndServe(":8080", nil)
}

// http.Handle - функция, которая регистрирует обработчик http.Handler
// http.HandlerFunc - адаптер. Тип данных, кот.явл. реализацией интерфейса http.Handler
