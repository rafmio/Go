// https://golang.howtos.io/using-the-io-reader-and-io-writer-interfaces-in-go/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
	}
	defer file.Close()

	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err.Error())
	} else {
		fmt.Printf("%d bytes has been read\n", n)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s", string(data[i]))
	}

	fmt.Println()
}
