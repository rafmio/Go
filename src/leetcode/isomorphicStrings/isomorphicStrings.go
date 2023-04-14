package isomorphic

import "fmt"

func isomorphicStrings(s string, t string) bool {
	mapChars := make(map[string]string)

	var result bool

	for i, value := range s {
		chr, ok := mapChars[string(value)] // проверяем наличие буквы в карте
		if ok {                            // если есть, то проверяем мэпится ли слово1-буква и слово2-буква
			fmt.Println("value", chr, "is in map")
			if string(chr) == string(t[i]) { // если слово1-буква == слово2-буква
				fmt.Println("chr == t[i]", chr, t[i])
				result = true // устанавливаем результат в тру
				continue      // продолжаем перебор
			} else { // если слово1-буква != слово1-буква
				fmt.Println("chr != t[i]", chr, t[i])
				result = false // результат - фолс
				break          // прерываем перебор
			}
		} else { // если текущей буква в перебираемом слове нет в карте...
			fmt.Println("there is no char in the map, adding...")
			mapChars[string(value)] = string(chr) // добавляем в карту
		}
	}

	fmt.Println(mapChars)

	return result
}
