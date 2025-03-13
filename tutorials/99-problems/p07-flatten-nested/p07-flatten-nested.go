package flatten

func Flatten(list [][]string) []string {
	result := make([]string, 0)

	for _, l := range list {
		for _, el := range l {
			result = append(result, string(el))
		}
	}

	return result
}
