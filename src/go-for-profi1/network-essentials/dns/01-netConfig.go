package main

import (
	"fmt"
	"net"
)

func main() {
	// net.Interfaces() возвращает все интерфейсы текущего компа в виде среза
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("ERROR making net.Interfaces():", err.Error())
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)

		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println("ERROR making net.InterfaceByName():", err.Error())
		}

		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}

		fmt.Println()
	}

}

// we may use '$ ip addr' command at command line
