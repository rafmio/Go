package dropelem

import (
	"reflect"
	"testing"
)

type TstCase struct {
	input    []string
	nThElem  int
	expected []string
	err      error
}

// test function for DropElem() than drops every N'th element from a list
func TestDropElem(t *testing.T) {
	tstCases := []TstCase{
		{
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
			nThElem:  2,
			expected: []string{"a", "c", "e", "g", "i", "k", "m", "o", "q", "s", "u", "w", "y"},
			err:      nil,
		},
		{
			input:    []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			nThElem:  3,
			expected: []string{"1", "2", "4", "5", "7", "8", "10"},
			err:      nil,
		},
		{
			input:    []string{"a", "b"},
			nThElem:  0,
			expected: []string{"a", "b"},
			err:      nil,
		},
		{
			input:    []string{"a", "b", "c"},
			nThElem:  -1,
			expected: []string{"a", "b", "c"},
			err:      ErrNegativeIdx,
		},
		{
			input:    []string{"a", "b", "c"},
			nThElem:  1,
			expected: []string{},
			err:      nil,
		},
	}

	for _, tstCase := range tstCases {
		got, err := DropElem(tstCase.input, tstCase.nThElem)

		if !reflect.DeepEqual(got, tstCase.expected) {
			t.Errorf("DropElem() = %v, want %v", got, tstCase.expected)
		}

		if err != tstCase.err {
			t.Errorf("DropElem() error = %v, want %v", err, tstCase.err)
		}
	}
}
