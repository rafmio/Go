package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("opening the file: ", err.Error())
	}
	defer file.Close()

	reader := io.Reader(file)
	minBuffer := make([]byte, 1)
	for {
		n, err := reader.Read(minBuffer)

		fmt.Printf("%v %v %v\n", string(minBuffer), n, err)

		if err != nil {
			break
		}
	}

	
}