package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file2.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write([]byte("Hello-mello, Tosy-Bosy"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Exit(0)
}

// func (b *Writer) Write(p []byte) (nn int, err error)
// writes the contents of p into buffer
// returns number of bytes written
// if  nn < len(p) - returns an error explaining why the write is short

// func (b *Writer) Flush() error
// writes any buffered data to the underlying io.Writer
