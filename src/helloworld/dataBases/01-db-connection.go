package main

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
)

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "host")
}
