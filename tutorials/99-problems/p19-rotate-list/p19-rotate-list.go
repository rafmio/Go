package rotatelist

import "errors"

var (
	ErrListIsEmpty = errors.New("the input list is empty")
	ErrInvalidN    = errors.New("the N is negative or out of range")
)

func RotateList(list []string, N int) ([]string, error) {
	result := make([]string, 0, len(list))
	if len(list) == 0 {
		return result, ErrListIsEmpty
	}
	if N < 0 || N > len(list) {
		return result, ErrInvalidN
	}

	idx := len(list)
	counter := 0

	for i := N; i > 0; i-- {
		idx--
		if idx < 0 {
			idx = len(list) - 1
		}
		counter++
	}

	plase := len(list) - idx

	result = append(result, list[plase:]...)
	result = append(result, list[:plase]...)

	return result, nil
}
