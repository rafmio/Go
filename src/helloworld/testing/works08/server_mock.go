package client

import (
	"fmt"
	"net/http"
	"time"
)

func runServer() {
	mux := http.NewServeMux()
	currentDate := time.Now()
	dateDDMMYYYY := currentDate.Format("02-01-2006")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, dateDDMMYYYY)
	})

	go http.ListenAndServe(":8080", mux)

	_ = runClient()
	// ret := runClient()
	// fmt.Println("ret: ", ret)
}
