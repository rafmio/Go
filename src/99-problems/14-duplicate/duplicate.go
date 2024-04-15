package duplicate

func Duplicate(list []string) []string {

	result := make([]string, 0)
	if len(list) == 0 {
		return result
	}

	for i := 0; i < len(list); i++ {
		result = append(result, list[i])
		result = append(result, list[i])
	}

	return result
}
