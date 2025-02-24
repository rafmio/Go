// https://www.kelche.co/blog/go/golang-bufio/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file-mile2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	nw, err := writer.Write([]byte("Hello-mello Huggy Wuggy\n"))
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Printf("%d bytes has been wrote\n", nw)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// The default buffer size for a buffered writer is 4096 bytes. This means that
// the data will be written to the writer in chuncks of 4096 bytes.
// For changing default size of buffer you should use bufio.NewWriterSize()
