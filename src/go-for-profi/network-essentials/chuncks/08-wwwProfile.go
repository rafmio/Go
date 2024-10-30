package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number:", PORT)
	} else {
		// PORT = ":" + arguments[1]
		PORT = arguments[1]
		fmt.Println("Using provided port number:", PORT)
	}

	r := http.NewServeMux()
	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/", myHandler)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}

// $ go run 08-wwwProfile.go
// $ go tool pprof http://localhost:8001/debug/pprof/profile

// type in the browser bar:
// http://localhost:8001/debug/pprof/
