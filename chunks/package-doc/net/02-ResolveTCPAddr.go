package main

import (
	"fmt"
	"net"
)

func resolve(host, port string) {
	address := net.JoinHostPort(host, port)

	// resolve TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println("Error resolving address:", err)
	}
	fmt.Println("Resoved TCP Address:", tcpAddr)
}

func main() {
	resolve("localhost", "8080")
	resolve("192.168.0.1", "9080")
	resolve("google.com", "10080")
	resolve("raf", "30080")
	resolve("194.58.102.129", "5432")

	fmt.Println("End program")
}
