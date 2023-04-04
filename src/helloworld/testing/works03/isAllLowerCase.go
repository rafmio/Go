package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func IsAllLowerCase(input string) (bool, error) {
	re, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		return false, err
	}

	loc := re.FindStringIndex(input)
	if len(loc) > 0 {
		return false, fmt.Errorf("input must contain only characters")
	}

	for _, ch := range input {
		if unicode.IsUpper(ch) {
			return false, nil
		}
	}

	return true, nil
}
