package insertelem

func InsertElement(list []string, position int, element string) []string {
	result := make([]string, 0, len(list)+1)

	if position <= 0 {
		result = append(result, element)
		result = append(result, list...)
		return result
	} else if position >= len(list) {
		result = append(result, list...)
		result = append(result, element)
		return result
	}

	result = append(result, list[:position-1]...)
	result = append(result, element)
	result = append(result, list[position-1:]...)

	return result
}
