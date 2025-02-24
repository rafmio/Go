package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter a number between 1 and 99: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if num < 1 || num > 99 {
			fmt.Println("Number should be between 1 and 99.")
			continue
		}

		result := num * 10
		fmt.Printf("%d * 10 = %d\n", num, result)
		break
	}
}
