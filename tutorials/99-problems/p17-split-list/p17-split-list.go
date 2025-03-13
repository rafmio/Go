package splitlist

import "errors"

var ErrOutOfRange = errors.New("length is out of range")
var ErrNegativeLen = errors.New("length is negative")

func SplitList(list []string, lenFirstPart int) ([][]string, error) {
	result := make([][]string, 0)

	if lenFirstPart < 0 {
		return result, ErrNegativeLen
	} else if lenFirstPart == 0 {
		result = append(result, list)
		return result, nil
	} else if lenFirstPart > len(list) {
		return result, ErrOutOfRange
	}

	result = append(result, list[:lenFirstPart])
	result = append(result, list[lenFirstPart:])

	return result, nil
}
