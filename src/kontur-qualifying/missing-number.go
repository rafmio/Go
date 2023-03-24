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

	desiredValue := mapValues(numSlice)
	fmt.Println(desiredValue)
}

func fillSlice(num int, numSlice []int) {
	reader := bufio.NewReader(os.Stdin)
	var elem int
	for i := 0; i < num; i++ {
		fmt.Fscan(reader, &elem)
		numSlice[i] = elem
	}
}

func mapValues(numSlice []int) int {
	var desiredValue int
	var valuesPairs = make(map[int]int)
	for _, val := range numSlice {
		valuesPairs[val] += 1
	}

	for key, value := range valuesPairs {
		fmt.Println(key, ": ", value)
	}

	for key, value := range valuesPairs {
		if value == 1 {
			desiredValue = key
		}
	}

	return desiredValue
}
