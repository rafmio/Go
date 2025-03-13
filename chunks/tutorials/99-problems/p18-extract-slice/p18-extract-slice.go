package extractslice

import (
	"errors"
)

var (
	ErrIncorrectSls     = errors.New("the parameters for extracting are incorrect")
	ErrInputListIsEmpty = errors.New("the input list is empty")
	ErrOutOfRange       = errors.New("the parameters for extracting are out of range or negative")
	ErrInputListIsNil   = errors.New("the input list is nil")
)

func ExtractSlice(list []string, sls [2]int) ([]string, error) {

	result := make([]string, 0)

	if len(list) == 0 {
		return result, ErrInputListIsEmpty
	} else if sls[0] < 1 || sls[1] < 1 {
		return result, ErrOutOfRange
	} else if sls[1] > len(list) {
		return result, ErrOutOfRange
	} else if sls[0] > sls[1] {
		return result, ErrIncorrectSls
	}

	result = append(result, list[(sls[0]-1):(sls[1])]...)

	return result, nil
}
