// https://golang.howtos.io/understanding-the-bufio-scanner-in-go/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
//	scanner := bufio.NewScanner(os.Stdin)
	file, _ := os.Open("03-bufio.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Scan lines until there is no more input
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Scanned line:", line)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning:", err)
	}
}
