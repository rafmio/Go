package isomorphic

import (
	"testing"
)

type tstCase struct {
	str1   string
	str2   string
	result bool
}

type tstCases []tstCase

func TestIsomorphicStrings(t *testing.T) {
	TableTstCases := tstCases{
		{str1: "egg", str2: "add", result: true},
		{str1: "foo", str2: "bar", result: false},
		{str1: "paper", str2: "title", result: true},
		{str1: "aab", str2: "xxy", result: true},
		{str1: "aab", str2: "xyz", result: false},
		{str1: "acab", str2: "xcxy", result: true},
	}

	for _, tst := range TableTstCases {
		got := isomorphicStrings(tst.str1, tst.str2)
		want := tst.result
		if got != want {
			t.Errorf("str1: %v, str2: %v -- got %v, want %v", tst.str1, tst.str2, got, want)
		}
	}
}
