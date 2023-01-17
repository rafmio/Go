package main

import "net/http"

func init() {
  http.HandleFunc("/redirect1", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/redirect2", http.StatusTemporaryRedirect)
  })
  http.HandleFunc("/redirect2", func(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/redirect1", http.StatusTemporaryRedirect)
  })
}
