package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func getIntInput() int {
	var input string
	var num int
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an integer: ")

		input, _ = reader.ReadString('\n')
		input = input[:len(input)-1]

		parsedNum, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: incorrect input, try again")
			continue
		}

		num = parsedNum

		break
	}

	return num
}

func makeAndFillOneDslice(num int) []int {
	sls := make([]int, num)

	for i := range sls {
		sls[i] = rand.Intn(num * 10)
	}

	return sls
}

func makeAndFillTwoDslice(num int) [][]int {
	rows := rand.Intn(num)
	cols := rand.Intn(num)

	if rows < 1 {
		rows = 1
	}

	if cols < 1 {
		cols = 1
	}

	sls := make([][]int, rows)
	for i := range sls {
		sls[i] = makeAndFillOneDslice(cols)
	}

	return sls
}

func main() {
	num := getIntInput()
	fmt.Println("You've entered: ", num)

	// sls := makeAndFillOneDslice(num)

	// for i, val := range sls {
	// 	fmt.Println(i, ":", val)
	// }

	twoDsls := makeAndFillTwoDslice(num)

	for _, val := range twoDsls {
		for _, value := range val {
			fmt.Printf("%d\t", value)
		}
		fmt.Println("\n")
	}
}
