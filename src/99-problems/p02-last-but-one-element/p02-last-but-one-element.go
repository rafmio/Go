package lastbutone

import (
	"errors"
)

// There are on 'exceptions' in Go
var ErrNoSuchElementException = errors.New("No such element exception")

// I can't pass arbitrary value type, now done just for int
func LastButOne(sls []int) (int, error) {
	if len(sls) == 0 || len(sls) == 1 {
		return 0, ErrNoSuchElementException
	}

	return sls[len(sls)-2], nil
}
