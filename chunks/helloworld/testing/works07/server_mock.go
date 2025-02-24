package client

import (
	"fmt"
	"net/http"
	"time"
)

func runServer() {
	mux := http.NewServeMux()

	currentDate := time.Now()
	dateDDMMYYYY := currentDate.Format("01-02-2006")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, dateDDMMYYYY)
	})

	http.ListenAndServe(":8080", mux)
}
