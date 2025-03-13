package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error fetching network interfaces:", err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("Interface Flags:", i.Flags.String())
		fmt.Println("Interface MTU:", i.MTU)
		fmt.Println("Interface HW Address:", i.HardwareAddr)
		fmt.Println("---------------------------------------")
	}
}
