// https://golang.howtos.io/using-the-io-reader-and-io-writer-interfaces-in-go/
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
	} else {
		fmt.Println("The file has been opened")
	}

	defer file.Close()

	data := []byte("Hello, World\n")
	nw, err := file.Write(data)
	if err != nil {
		fmt.Println("Error writing file:", err.Error())
	} else {
		fmt.Printf("%d bytes has been wrote\n", nw)
	}

}
