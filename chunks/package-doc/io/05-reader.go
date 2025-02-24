// https://youtu.be/O-MeKOuvzYE?si=LQGlMwTnz2ESBvrZ
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileReader()
}

func fileReader() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
	}
	defer f.Close()

	buf := make([]byte, 30) // you may set 20, 5, 40 etc instead of '10'
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file:", err.Error())
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
