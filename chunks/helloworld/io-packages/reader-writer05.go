package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("opening the file:", err.Error())
	}
	defer file.Close()

	reader := io.Reader(file)
	buffer, err := io.ReadAll(reader)
	fmt.Printf("RreadAll buffer={%v}, err={%v}\n", string(buffer), err)
}