package kids

import (
	"fmt"
	"testing"
)

type tstCase struct {
	candies      []int
	extraCandies int
	answer       []bool
}

func TestKidsWithCandies(t *testing.T) {
	TestTable := []tstCase{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for idx, val := range TestTable {
		got := kidsWithCandies(val.candies, val.extraCandies)

		for i := 0; i < len(got); i++ {
			if got[i] != val.answer[i] {
				t.Errorf("Test case: %d, subcase: %d got %t want %t", idx, i, got[i], val.answer[i])
			} else if got[i] == val.answer[i] {
				fmt.Printf("case: %d, subcase: %d, got: %t, want %t\n", idx, i, got[i], val.answer[i])
			}
		}
	}
}
