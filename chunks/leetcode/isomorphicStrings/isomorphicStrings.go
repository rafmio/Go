package isomorphic

func isomorphicStrings(s string, t string) bool {
	sPattern, tPattern := map[uint8]int{}, map[uint8]int{}
	for idx := range s {
		if sPattern[s[idx]] != tPattern[t[idx]] {
			return false
		} else {
			sPattern[s[idx]] = idx + 1
			tPattern[t[idx]] = idx + 1
		}
	}

	return true
}
