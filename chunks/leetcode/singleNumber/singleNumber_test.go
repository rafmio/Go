package singleNum

import (
	"testing"
)

type TstCase struct {
	nums   []int
	answer int
}

type TstCases []TstCase

func TestSingleNumber(t *testing.T) {
	casesTable := TstCases{
		{nums: []int{2, 2, 1}, answer: 1},
		{nums: []int{4, 1, 2, 1, 2}, answer: 4},
		{nums: []int{1}, answer: 1},
	}

	for _, tstCase := range casesTable {
		got := singleNumber(tstCase.nums)
		want := tstCase.answer
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
}
