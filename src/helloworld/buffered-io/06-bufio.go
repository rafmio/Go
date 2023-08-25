// https://www.kelche.co/blog/go/golang-bufio/
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	EXIT_SUCCESS = 0
	EXIT_FAILURE = 1
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Too few arguments. Usage: '$ %s filepath\n'", os.Args[0])
		os.Exit(EXIT_FAILURE)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("opening file: ", err.Error())
		os.Exit(EXIT_FAILURE)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("bufio.ReadString():", err.Error())
		os.Exit(EXIT_FAILURE)
	}

	fmt.Println(data)

	os.Exit(EXIT_SUCCESS)
}
