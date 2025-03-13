package problems

func FindLastElement(list []string) string {
	if len(list) == 0 {
		return ""
	} else {
		return list[len(list)-1]
	}
}
