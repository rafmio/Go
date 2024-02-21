package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error net.Interfaces():", err.Error())
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Name: %v\n", i.Name)
		fmt.Println("Interface Flags:", i.Flags.String())
		fmt.Println("Interface MTU:", i.MTU)
		fmt.Println("Interface Hardware Address:", i.HardwareAddr)
		fmt.Println()
	}
}

// Interface struct:
// type Interface struct {
// 	Index        int
// 	MTU          int
// 	Name         string
// 	HardwareAddr HardwareAddr
// 	Flags        Flags
// }
