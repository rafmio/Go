package handlers

import (
	"os"
)

func OpenAndReadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	result := make([]string, 0)

	// fill result slice

	return result, nil
}
