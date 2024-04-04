package reverseList

import "errors"

var ErrListIsEmpty = errors.New("list is empty")

func ReverseList(list []int) ([]int, error) {

	// Easy way: use strlib function:
	// slices.Reverse(list)
	// return list, nil

	// More complicated way:
	if len(list) == 0 {
		return nil, ErrListIsEmpty
	} else {
		for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
			list[i], list[j] = list[j], list[i]
		}
	}

	return list, nil
}
