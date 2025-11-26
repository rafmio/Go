package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	helloHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	expected := "Hello, World!"
	if w.Body.String() != expected {
		t.Errorf("Expected body %q, got %q", expected, w.Body.String())
	}
}

func TestTimeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/time", nil)
	w := httptest.NewRecorder()

	timeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Current time: " + time.Now().Format(time.RFC1123)))
	})
	timeHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	body := w.Body.String()
	if !stringContains(body, "Current time: ") {
		t.Errorf("Expected body to contain 'Current time: ', got %q", body)
	}
}

func TestCountHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/count", nil)
	w := httptest.NewRecorder()

	countHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Count: 42"))
	})
	countHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", w.Code)
	}

	expected := "Count: 42"
	if w.Body.String() != expected {
		t.Errorf("Expected body %q, got %q", expected, w.Body.String())
	}
}

// Вспомогательная функция
func stringContains(s, substr string) bool {
	return len(substr) <= len(s) && s[:len(substr)] == substr
}
