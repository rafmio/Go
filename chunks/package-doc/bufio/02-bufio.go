package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan lines until there is no more input
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Scanned line:", line)
		if line == "exit" {
			os.Exit(0)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}
}
