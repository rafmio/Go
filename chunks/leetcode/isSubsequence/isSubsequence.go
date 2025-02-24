package subsequence

func isSubsequence(s string, t string) bool {
	sRune := []rune(s)
	tRune := []rune(t)
	var result bool
	var sPosition int

	if len(s) == 0 {
		return true
	} else if len(s) == 0 && len(t) == 0 {
		return true
	} else if len(s) > len(t) {
		return false
	}

	for _, letter := range tRune {

		if sPosition == len(sRune) {
			break
		} else if letter == sRune[sPosition] {
			sPosition++
			result = true
			continue
		} else {
			result = false
			continue
		}

	}

	if sPosition < len(sRune) {
		result = false
	}

	return result
}
