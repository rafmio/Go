package modEnc

func RunLengthEncoding(list []string) map[string]int {
	result := make(map[string]int)

	for _, letter := range list {
		if result[letter] == 0 {
			result[letter] = 1
		} else {
			result[letter] += 1
		}
	}

	for key, val := range result {
		if val > 1 {
			delete(result, key)
		}
	}

	return result
}
