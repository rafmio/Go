package remove

import "errors"

var (
	ErrEmptyInput   = errors.New("input list is empty")
	ErrKIsIncorrect = errors.New("K is out of range or negative")
)

func Remove(list []string, K int) ([]string, error) {
	if len(list) == 0 {
		return []string{}, ErrEmptyInput
	}

	if K < 0 || K > len(list) {
		return []string{}, ErrKIsIncorrect
	}

	list = append(list[:K], list[K+1:]...)

	return list, nil
}
