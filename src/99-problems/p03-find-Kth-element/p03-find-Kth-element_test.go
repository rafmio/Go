package kthelement

import (
	"testing"
)

func TestFindKthElement(t *testing.T) {

	type Cases struct {
		kthElem int
		slice   []int
		want    int
		errWant error
	}

	cases := []Cases{
		{kthElem: 3, slice: []int{1, 2, 3, 4, 5}, want: 4, errWant: nil},
		{kthElem: 5, slice: []int{34, 43, 55, 3403, 1, 300, 3, 9}, want: 300, errWant: nil},
		{kthElem: 4, slice: []int{3, 4, 5, 11}, want: 0, errWant: ErrKOutOfRange},
	}

	t.Run("finds Kth element of the list if list not empty", func(t *testing.T) {
		for _, cs := range cases {
			got, err := FindKthElement(cs.slice, cs.kthElem)
			want := cs.want
			wantError := cs.errWant

			if err != wantError {
				t.Errorf("got %d want %d", err, wantError)
			}

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		}
	})

}
