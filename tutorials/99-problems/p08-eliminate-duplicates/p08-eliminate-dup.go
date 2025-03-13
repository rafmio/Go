package eldup

func EliminateDup(sls []string) []string {
	var checkedElem string

	for i := 0; i < len(sls); i++ {
		if checkedElem != sls[i] {
			checkedElem = sls[i]
		} else {
			sls = append(sls[:i], sls[i+1:]...)
			i--
		}
	}

	return sls
}

func EliminateDupGenericComparable[T comparable](sls []T) []T {
	var checkedElem T

	for i := 0; i < len(sls); i++ {
		if checkedElem != sls[i] {
			checkedElem = sls[i]
		} else {
			sls = append(sls[:i], sls[i+1:]...)
			i--
		}
	}

	return sls
}
