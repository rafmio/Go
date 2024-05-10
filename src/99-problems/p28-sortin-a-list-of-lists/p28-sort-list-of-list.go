package sortlistoflist

import "errors"

var ErrEmptyList = errors.New("the list is empty")

func SortListOfList(list [][]string) ([][]string, error) {
	if len(list) == 0 {
		return [][]string{}, ErrEmptyList
	} else if len(list) == 1 {
		return list, nil
	}

	var isOrdered bool

	for !isOrdered {
		for i := 0; i < len(list); i++ {
			if i < len(list)-1 {
				if len(list[i]) > len(list[i+1]) {
					list[i], list[i+1] = list[i+1], list[i]
					i = 0
					isOrdered = false
				} else {
					isOrdered = true
				}
			}
		}

	}

	return list, nil
}
