// https://www.kelche.co/blog/go/golang-bufio/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("error:", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	data := make([]byte, 128)
	_, err = reader.Read(data)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(data))
}
