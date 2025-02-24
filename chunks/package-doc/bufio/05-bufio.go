// https://metanit.com/go/tutorial/8.9.php
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("some.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err.Error())
		return
	} else {
		fmt.Println("the file was opened")
	}
	defer file.Close()

	// bufio.NewReader returns *Reader
	// Reader - то, откуда читаем
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err.Error())
				return
			}
		}
		fmt.Print(line)
	}
}
