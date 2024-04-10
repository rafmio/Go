package packdup

// PackDupStr accepts a string slice ([]string) pack duplicates into it's own
// slice and do similar with unique elements - pack them into own slice with
// single one element

func PackDupStr(list []string) [][]string {
	var result [][]string
	// currentGroup := []string{}
	currentGroup := make([]string, 0)

	for i, letter := range list {
		if i == 0 || list[i-1] != letter { // TODO: проверить не будет ли out of range при list[i-1]
			// Если текущая буква отличается от предыдущей или это первый элемент, начинаем новый слайс
			currentGroup = []string{letter}
		} else {
			// Добавляем букву к текущему слайсу
			currentGroup = append(currentGroup, letter)
		}

		// Добавляем текущий слайс в результат, если мы достигли конца списка или следующая буква другая
		if i+1 >= len(list) || list[i+1] != letter {
			result = append(result, currentGroup)
		}
	}

	return result
}
