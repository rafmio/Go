package majority

import (
	"testing"
)

type inputCase struct {
	nums   []int
	answer int
}

func TestMajorityElement(t *testing.T) {
	TestTable := []inputCase{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{-1, 2, -1, 3, -1, 10, -1}, -1},
	}

	for _, tstCase := range TestTable {
		got := majorityElement(tstCase.nums)
		want := tstCase.answer

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
}
