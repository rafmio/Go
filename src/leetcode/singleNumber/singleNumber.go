package singleNum

func singleNumber(nums []int) int {
	var result int
	checkAppearence := make(map[int]int)
	for _, value := range nums {
		checkAppearence[value]++
	}

	for key, value := range checkAppearence {
		if value == 1 {
			result = key
		}
	}

	return result
}
