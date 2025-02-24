package main

import (
	"fmt"
)

func PackConsecutiveDuplicates(input []string) [][]string {
	result := make([][]string, 0)

	for _, value := range input {
		if len(result) == 0 {
			result = append(result, []string{value})
		} else {
			lastIndex := len(result) - 1
			lastGroup := result[lastIndex]
			if lastGroup[0] == value {
				result[lastIndex] = append(lastGroup, value)
			} else {
				result = append(result, []string{value})
			}
		}
	}

	return result
}

func main() {
	input := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}
	result := PackConsecutiveDuplicates(input)
	fmt.Println(result)
}
