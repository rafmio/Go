package splitlist

import (
	"reflect"
	"testing"
)

type TstCase struct {
	name         string
	input        []string
	lenFirstPart int
	expected     [][]string
	err          error
}

func TestSplitList(t *testing.T) {
	tstCases := []TstCase{
		{
			name:         "length 5",
			input:        []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			lenFirstPart: 5,
			expected: [][]string{
				{"a", "b", "c", "d", "e"},
				{"f", "g", "h", "i", "j"},
			},
			err: nil,
		},
		{
			name:         "length 1",
			input:        []string{"A", "B", "C"},
			lenFirstPart: 1,
			expected: [][]string{
				{"A"},
				{"B", "C"},
			},
			err: nil,
		},
		{
			name:         "length 3",
			input:        []string{"AA", "BB", "CC", "DD"},
			lenFirstPart: 3,
			expected: [][]string{
				{"AA", "BB", "CC"},
				{"DD"},
			},
			err: nil,
		},
		{
			name:         "zero length",
			input:        []string{"X", "Y", "Z"},
			lenFirstPart: 0,
			expected: [][]string{
				{"X", "Y", "Z"},
			},
			err: nil,
		},
		{
			name:         "the length is out of range",
			input:        []string{"M", "N", "O", "P", "Q"},
			lenFirstPart: 10,
			expected:     [][]string{},
			err:          ErrOutOfRange,
		},
		{
			name:         "the length is negative",
			input:        []string{"M", "N", "O", "P", "Q"},
			lenFirstPart: -2,
			expected:     [][]string{},
			err:          ErrNegativeLen,
		},
	}

	for _, tstCase := range tstCases {
		t.Run(tstCase.name, func(t *testing.T) {
			got, err := SplitList(tstCase.input, tstCase.lenFirstPart)

			if !reflect.DeepEqual(got, tstCase.expected) {
				t.Errorf("SplitList(): got %v, want: %v", got, tstCase.expected)
			}

			if err != tstCase.err {
				t.Errorf("SplitList(): got %v, want: %v", err, tstCase.err)
			}
		})
	}
}
