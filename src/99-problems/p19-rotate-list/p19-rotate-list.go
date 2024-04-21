package rotatelist

import "errors"

var (
	ErrListIsEmpty = errors.New("the input list is empty")
	ErrInvalidN    = errors.New("the N is negative or out of range")
)

func RotateList(list []string, N int) ([]string, error) {

	result := make([]string, 0, len(list))

	if len(list) == 0 {
		return list, ErrListIsEmpty
	}

	if N < 0 || N > len(list) {
		return result, ErrInvalidN
	}

	if N == len(list) {
		return list, nil
	}

	if N == 0 {
		return list, nil
	}

	idx := len(list) - 1

	// find index
	for N > 0 {
		idx--
		if idx < 0 {
			idx = len(list) - 1
		} else {
			N--
		}
	}

	for i := 0; i < len(list); i++ {
		// result[i] = list[idx]
		result = append(result, list[idx])
		idx++
		if idx >= len(list) {
			idx = 0
		}
	}

	return result, nil
}
