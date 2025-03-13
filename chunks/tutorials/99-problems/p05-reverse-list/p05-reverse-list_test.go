package reverseList

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type ListsForReverse struct {
		originList    []int
		originListLen int
		reversedList  []int
	}

	listCases := []ListsForReverse{
		{originList: []int{1}, originListLen: 1, reversedList: []int{1}},
		{originList: []int{1, 2}, originListLen: 2, reversedList: []int{2, 1}},
		{originList: []int{1, 2, 3, 4, 5}, originListLen: 5, reversedList: []int{5, 4, 3, 2, 1}},
		{originList: []int{1, 2, 3, 4, 5, 6}, originListLen: 6, reversedList: []int{6, 5, 4, 3, 2, 1}},
		{originList: []int{1, 2, 3, 4, 5, 6, 7}, originListLen: 7, reversedList: []int{7, 6, 5, 4, 3, 2, 1}},
		{originList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, originListLen: 10, reversedList: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	t.Run("reverses a list", func(t *testing.T) {

		for _, listCase := range listCases {
			got, _ := ReverseList(listCase.originList)

			if !reflect.DeepEqual(got, listCase.reversedList) {
				t.Errorf("list was reversed wrong")
			}

			if len(got) != listCase.originListLen {
				t.Errorf("len of got list: %d, len origin list: %d", len(got), listCase.originListLen)
			}
		}

		// TODO: check errors when slice is empty

	})

	t.Run("when empty slice was passed", func(t *testing.T) {
		_, err := ReverseList([]int{})
		if err != ErrListIsEmpty {
			t.Errorf("got %v, want %v", err, ErrListIsEmpty)
		}
	})
}
