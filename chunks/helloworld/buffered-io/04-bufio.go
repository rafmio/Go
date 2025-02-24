// https://www.kelche.co/blog/go/golang-bufio/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Too few arguments")
	}
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0)
	if err != nil {
		log.Fatal("opening file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data := make([]byte, 100)
	_, err = reader.Read(data)
	if err != nil {
		log.Fatal("reading file:", err.Error())
	}

	fmt.Println(string(data))
}
