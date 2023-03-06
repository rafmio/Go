package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file *os.File
	file = os.Stdin
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		enteredVal := scanner.Text()
		fmt.Println(">", enteredVal)
		if enteredVal == string("END") {
			break
		}
	}
}
