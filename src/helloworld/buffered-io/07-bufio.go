// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	EXIT_SUCCESS = 0
	EXIT_FAILURE = 1
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s filename\n", os.Args[0])
		os.Exit(EXIT_FAILURE)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("opening file:", err.Error())
		os.Exit(EXIT_FAILURE)
	}
	defer file.Close()

	var counter int = 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(counter, " ", scanner.Text())
		counter++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("scanner.Err():", err.Error())
	}

	os.Exit(EXIT_SUCCESS)
}
