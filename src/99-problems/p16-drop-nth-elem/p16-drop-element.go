package dropelem

import (
	"errors"
)

var ErrNegativeIdx = errors.New("the index is negative")

// the function drops every Nth element (elment with elemIdx)
func DropElem(list []string, elemIdx int) ([]string, error) {
	if elemIdx < 0 {
		return list, ErrNegativeIdx
	} else if elemIdx == 0 {
		return list, nil
	} else if elemIdx == 1 {
		emptySls := make([]string, 0)
		return emptySls, nil
	}

	return list, nil
}
