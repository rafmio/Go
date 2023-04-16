package subsequence

import (
	"testing"
)

type tstCase struct {
	s      string
	t      string
	result bool
}

type tstCases []tstCase

func TestIsSubsequence(t *testing.T) {
	cases := tstCases{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"", "ahbgdc", true},
		{"", "", true},
		{"acb", "ahbgdc", false},
		{"b", "abc", true},
	}

	for _, tstCs := range cases {
		got := isSubsequence(tstCs.s, tstCs.t)
		want := tstCs.result
		if got != want {
			t.Errorf("s: %q, t: %q, got %t want %t", tstCs.s, tstCs.t, got, want)
		}
	}
}
