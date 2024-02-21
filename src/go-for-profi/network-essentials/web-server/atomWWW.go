package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"
)

var count int32

func handleAll(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count, 1)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&count)
	fmt.Println("Count:", temp)
	fmt.Fprintf(w, "<h1 align=\"center\">%d</h1>", count)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	http.HandleFunc("/getCounter", getCounter)
	http.HandleFunc("/", handleAll)

	http.ListenAndServe(":8080", nil)
}

// ab - Apache HTTP server benchmarking tool
// $ sudo apt install apache2-utils
// $ ab -n 1500 -c 100 http://localhost:8080/
// Показанная выше команда ab(1) отправляет 1500 запросов, при этом количество
// одновременных запросов равно 100
