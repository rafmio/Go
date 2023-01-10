package main

import (
  "net/http"
  "encoding/json"
)

func HandleJsonRequest(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(Products)
}

func init() {
  http.HandleFunc("/json", HandleJsonRequest)
}
// The init creates a route, which means that request for /json will be processed
// by the HandleJsonRequest - to encode slice of Product
