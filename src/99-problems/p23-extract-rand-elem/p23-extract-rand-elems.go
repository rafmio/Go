package extractrand

import (
	"errors"
	"math/rand"
)

var (
	ErrIncorrectNum   = errors.New("the num is negative or out of range")
	ErrEmptyInputList = errors.New("the input list is empty")
)

func ExtractRandElems(list []string, num int) ([]string, error) {
	result := make([]string, 0)

	if len(list) == 0 {
		return result, ErrEmptyInputList
	}

	if num < 0 || num > len(list) {
		return result, ErrIncorrectNum
	}

	for i := 0; i < num; i++ {
		idx := rand.Intn(num)
		result = append(result, list[idx])
	}

	return result, nil
}
