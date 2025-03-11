package main

import (
	"fmt"
	"net"
)

func main() {
	host := "localhost"
	port := "8080"

	// create address string
	address := net.JoinHostPort(host, port)

	// resolve TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("Error resolving address:", err)
	}

	fmt.Println("Resoved TCP Address:", tcpAddr)
}
