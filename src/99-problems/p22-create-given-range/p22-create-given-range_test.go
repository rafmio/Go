package givenrange

import (
	"reflect"
	"testing"
)

type TstCase struct {
	input [2]int
	want  []int
}

func TestGivenRange(t *testing.T) {
	tstCases := []TstCase{
		{
			input: [2]int{4, 9},
			want:  []int{4, 5, 6, 7, 8, 9},
		},
		{
			input: [2]int{1, 10},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			input: [2]int{10, 1},
			want:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			input: [2]int{1, 1},
			want:  []int{1},
		},
		{
			input: [2]int{-3, 3},
			want:  []int{-3, -2, -1, 0, 1, 2, 3},
		},
		{
			input: [2]int{3, -3},
			want:  []int{3, 2, 1, 0, -1, -2, -3},
		},
	}

	for _, tstCase := range tstCases {
		got := GivenRange(tstCase.input)
		want := tstCase.want

		if !reflect.DeepEqual(got, want) {
			t.Errorf("GivenRange(%v) = %v, want %v", tstCase.input, got, want)
		}
	}
}

func BenchmarkGivenRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GivenRange([2]int{1, 10})
		GivenRange([2]int{10, 1})
		GivenRange([2]int{4, 9})
		GivenRange([2]int{-3, 3})
		GivenRange([2]int{3, -3})
	}
}
