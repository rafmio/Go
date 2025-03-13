package givenrange

func GivenRange(list [2]int) []int {
	result := make([]int, 0)

	if list[0] == list[1] {
		result = append(result, list[0])
	} else if list[0] < list[1] {
		for i := list[0]; i <= list[1]; i++ {
			result = append(result, i)
		}
	} else {
		for i := list[0]; i >= list[1]; i-- {
			result = append(result, i)
		}
	}

	return result
}
