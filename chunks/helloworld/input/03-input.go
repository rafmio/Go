package main

import (
	"fmt"
	"math/rand"
)

func getUserInput() int {
	fmt.Printf("Please, guess the number: ")
	var input int
	fmt.Scan(&input)

	return input
}

func generateRandomNum() int {
	return rand.Intn(10) + 1
}

func compareNums(input, randomNum int) bool {
	right_message := "You are right! The answer is '%d'\n"

	var right_answer bool

	if input == randomNum {
		fmt.Printf(right_message, randomNum)
		right_answer = true
	} else if input > randomNum {
		fmt.Printf("Too high! My number is less than %d\n", input)
		right_answer = false
	} else {
		fmt.Printf("Too low! My number is greater than %d\n", input)
		right_answer = false
	}

	return right_answer
}

func main() {
	var input int
	randomNum := generateRandomNum()

	var right_answer bool

	for !right_answer {
		input = getUserInput()
		right_answer = compareNums(input, randomNum)

		if right_answer {
			break
		}

	}
}
