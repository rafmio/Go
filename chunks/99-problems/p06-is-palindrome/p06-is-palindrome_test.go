package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type casePlTst struct {
		input        []string
		isPalindrome bool
	}

	cases := []casePlTst{
		{input: []string{"a", "m", "m", "a"}, isPalindrome: true},
		{input: []string{"1", "2", "3", "2", "1"}, isPalindrome: true},
		{input: []string{"1", "2", "3", "2", "1", "2", "3", "2", "1"}, isPalindrome: true},
		{input: []string{"1", "2", "3", "2", "1", "2", "3", "2", "1", "2", "3", "2", "1"}, isPalindrome: true},
		{input: []string{"p", "a", "p", "a"}, isPalindrome: false},
		{input: []string{"1", "2", "3", "1"}, isPalindrome: false},
		{input: []string{"1", "2", "3", "4", "1"}, isPalindrome: false},
	}

	for i, c := range cases {
		got := IsPalindrome(c.input)
		if got != c.isPalindrome {
			t.Errorf("case: %d, got: %v, want: %v", i, got, c.isPalindrome)
		}
	}
}
