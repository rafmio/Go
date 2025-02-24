package title

func titleToNumber(columnTitle string) int {
	result := 0

	for _, char := range columnTitle {
		result *= 26
		result += int(char-'A') + 1
	}

	return result
}
