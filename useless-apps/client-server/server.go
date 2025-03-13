package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// http://numbersapi.com/%d

func numberInfoHandler(w http.ResponseWriter, r *http.Request) {
	// parse params
	params := r.URL.Query()
	numberStr := params.Get("number")

	if numberStr == "" {
		http.Error(w, "Number is required", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	numInfoAPIurl := "http://numbersapi.com/%d"
	readyUrl := fmt.Sprintf(numInfoAPIurl, number)

	response, err := http.Get(readyUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(body)
}

func main() {
	port := "9080"
	http.HandleFunc("/number-info", numberInfoHandler)
	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running...")
}
