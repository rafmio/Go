package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func scanInput(maxKey int) (int, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading input:", err)
			return 0, fmt.Errorf("error reading input: %w", err)
		}
		inputStr = strings.TrimSpace(inputStr)

		input, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if input < 1 || input > maxKey {
			fmt.Println("Number should be between 1 and", maxKey, ".")
			continue
		}

		return input, nil
	}
}

func main() {
	fmt.Printf("Please, enter a number: ")
	num, err := scanInput(10_000_000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You entered: %d\n", num)
}
