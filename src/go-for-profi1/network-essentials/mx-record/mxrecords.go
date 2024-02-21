package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := arguments[1]
	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("Error looking up mx-record:", err.Error())
		return
	}

	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}

// MX-записи указывают на почтовые серверы (mail server) домена

// $ host -t mx golang.com
