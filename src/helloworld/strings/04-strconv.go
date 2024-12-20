package main

import (
	"strconv"
	"strings"
)

func main() {
	fileNames := []string{
		"01-events.json",
		"02-products.json",
		"03-orders.json",
		"04-customers.json",
		"05-suppliers.json",
	}

	fileMap := make(map[int]string)

	for _, file := range fileNames {
		prefix := strings.Split(file, "-")[0]
		if len(prefix) != 2 || prefix[0] < '0' || prefix[0] > '9' || prefix[1] < '0' || prefix[1] > '9' {
			continue
		}

		key, err := strconv.Atoi(prefix)
		if err != nil {
			return
		}

		fileMap[key] = file
	}

	for key, file := range fileMap {
		println("File", key, "-", file)
	}
}
