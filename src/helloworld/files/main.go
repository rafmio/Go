package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <file> <message>")
		return
	}

	// f, err := os.Open(os.Args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer f.Close()

	f, err := os.OpenFile(os.Args[1], os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(os.Args[2])
	if err != nil {
		if err == os.ErrPermission {
			fmt.Println(err.Error())
		}
		fmt.Println(err.Error())
		return
	}
}
