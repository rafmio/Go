package duplicate

func Duplicate(list []string) []string {

	// result := make([]string, 0)
	result := make([]string, 0, len(list)*2)
	if len(list) == 0 {
		return result
	}

	// for i := 0; i < len(list); i++ {
	// 	result = append(result, list[i])
	// 	result = append(result, list[i])
	// }

	for _, value := range list {
		result = append(result, value, value)
	}

	return result
}
