package lastbutone

import (
	"errors"
)

var ErrNoSuchElementException = errors.New("No such element exception")

func LastButOne(sls []interface{}) (interface{}, error) {
	if len(sls) == 0 || len(sls) == 1 {
		return nil, ErrNoSuchElementException
	}

	return sls[len(sls)-2], nil
}
