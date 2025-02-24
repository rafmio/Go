package main

import (
	"strconv"
)

func Fooer(input int) string {
	isfoo := (input % 3) == 0
	if isfoo {
		return "Foo"
	}

	return strconv.Itoa(input)
}

// go test
// go test -v
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -bench=.
// go test -bench=FunctionName
