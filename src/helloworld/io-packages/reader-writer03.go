package main

import (
	"io"
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("creating the file:", err.Error())
	}
	defer file.Close()

	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello!"))
	if err != nil {
		fmt.Println("write to file: ", err.Error())
	}
	fmt.Println(n, err)
	fmt.Println()

	n, err = io.WriteString(writer, " Huggy-Wuggy")
	if err != nil {
		fmt.Println("WriteString(): ", err.Error())
	}
	fmt.Println(n, err)
}