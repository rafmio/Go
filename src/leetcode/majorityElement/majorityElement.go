package majority

func majorityElement(nums []int) int {
	amount := make(map[int]int)
	for _, value := range nums {
		amount[value]++
	}

	var mElem int
	var mAmount int

	for key, value := range amount {
		if mAmount == 0 {
			mAmount = value
			mElem = key
		} else if mAmount < value {
			mAmount = value
			mElem = key
		}
	}

	return mElem
}
