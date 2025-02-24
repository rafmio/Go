package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("opening file:", err.Error())
	} else {
		fmt.Println("file has been opened")
	}
	defer file.Close()

	data := []byte("Hihi-Haha, World!\n") // привели string к типу []byte
	nw, err := file.Write(data)
	if err != nil {
		fmt.Println("writing to file:", err.Error())
	} else {
		fmt.Printf("%d bytes has been wrote\n", nw)
	}

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

		fmt.Println(line)
	}
}

// 03-write.go
