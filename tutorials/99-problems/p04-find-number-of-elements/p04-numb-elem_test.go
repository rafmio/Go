package numelem

import (
	"testing"
)

func TestFindNumberOfElements(t *testing.T) {

	type List struct {
		list      []int
		numOfElem int
	}

	listsOfCases := []List{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
		{[]int{}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}, 26},
		{[]int{1, 2, 3}, 3},
	}

	t.Run("finds the number of elements of a list", func(t *testing.T) {

		for _, listOfCase := range listsOfCases {
			got := FindNumberOfElements(listOfCase.list)
			want := listOfCase.numOfElem

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		}

	})
}
