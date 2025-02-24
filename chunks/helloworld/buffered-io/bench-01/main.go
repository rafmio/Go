// https://www.kelche.co/blog/go/golang-bufio/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var filename string = "file.txt"

func funcToWithIO() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	for {
		_, err := file.Read(data)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func funcToWithBufio() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data := make([]byte, 100)
	for {
		_, err := reader.Read(data)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func createFile() {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	for i := 0; i < 1000000; i++ {
		file.Write([]byte("Hello-mello, Tosy-Bosy!"))
	}
}

func main() {
	createFile()
	funcToWithIO()
	funcToWithBufio()
}
