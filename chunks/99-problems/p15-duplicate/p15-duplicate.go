package duplicate

func Duplicate(times int, list []string) []string {
	if len(list) == 0 {
		return list
	}

	result := make([]string, 0)

	for i := 0; i < times; i++ {
		result = append(result, list...)
	}

	return result
}
