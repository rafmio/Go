package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num int
	fmt.Fscan(reader, &num)
	numSlice := make([]int, num)
	fillSlice(num, numSlice)

	var sum int
	for _, value := range numSlice {
		sum += value
	}

	var desiredValue int
	desiredValue = sum * -1
	fmt.Println(desiredValue)
}

func fillSlice(num int, numSlice []int) {
	// reader := bufio.NewReader(os.Stdin)
	var elem int
	for i := 0; i < num; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Fscan(reader, &elem)
		numSlice[i] = elem
	}
}
