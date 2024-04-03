package kthelement

import (
	"errors"
)

var ErrKOutOfRange = errors.New("the Kth element is out of range")

func FindKthElement(slice []int, kthElem int) (int, error) {

	if kthElem > len(slice)-1 || kthElem < 0 {
		return 0, ErrKOutOfRange
	}

	return slice[kthElem], nil

}
