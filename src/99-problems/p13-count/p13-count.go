package count

func Count(list []string) map[string]int {
	result := make(map[string]int)

	for _, element := range list {
		result[element]++
	}

	return result
}
