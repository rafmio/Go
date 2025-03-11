package main

import (
	"fmt"
	"net"
)

func main() {
	host := "localhost"
	port := "8080"

	// Join host and port
	address := net.JoinHostPort(host, port)
	fmt.Println("Address:", address)
}
