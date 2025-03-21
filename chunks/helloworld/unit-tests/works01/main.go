package main

import (
	"strconv"
	"strings"
)

const (
	kilometersToMiles = 0.621371
	milesToKiometers  = 1.60934
)

func Convert(from, to string) (string, error) {
	var result float64
	switch {
	case strings.HasSuffix(from, "mi"):
		miles, err := strconv.ParseFloat(from[:len(from)-2], 64)
		if err != nil {
			return "", err
		}
	}

	switch to {
	case "km":
		result = miles * milesToKiometers
	case "m":
		result = miles * milesToKiometers * 1000
	case "mi":
		result = miles
	}
}
