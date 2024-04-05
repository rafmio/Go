// https://www.kelche.co/blog/go/golang-bufio/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 128)
	nw, err := writer.Write([]byte("Hello-mello, Kissy Missy\n"))
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
