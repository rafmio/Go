package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("03-input.txt")
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Set the scanner split function to scan words
	scanner.Split(bufio.ScanWords)

	// Scan words until there is no more input
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println("Scanned word:", word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err.Error())
	}
}
