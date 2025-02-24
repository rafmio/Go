package kids

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	var greatest int
	for _, val := range candies {
		if greatest < val {
			greatest = val
		}
	}

	for i := 0; i < len(candies); i++ {
		if (candies[i] + extraCandies) > greatest {
			result[i] = true
		}
	}

	return result
}
