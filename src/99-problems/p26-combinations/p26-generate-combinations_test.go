package combinations

import (
	"testing"
)

type TstCase struct {
	name     string
	input    []string
	k        int
	expected [][]string
	err      error
}

// 'FindCombinations()' Generate the combinations of K distinct objects
// chosen from the N elements of a list
func TestFindCombinations(t *testing.T) {
	tstCases := []TstCase{
		{
			name:  "list of 2 elements, k == 1",
			input: []string{"a", "b"},
			k:     2,
			expected: [][]string{
				{"a", "b"},
				{"b", "a"},
			},
			err: nil,
		},
		{
			name:  "list of 3 elements, k == 2",
			input: []string{"a", "b", "c"},
			k:     2,
			expected: [][]string{
				{"a", "b"},
				{"a", "c"},
				{"b", "a"},
				{"b", "c"},
				{"c", "a"},
				{"c", "b"},
			},
		},
	}
}
