package extractslice

import (
	"reflect"
	"testing"
)

type TstCase struct {
	name     string
	input    []string
	sls      [2]int
	expected []string
	err      error
}

func TestExtractSlice(t *testing.T) {
	tstCases := []TstCase{
		{
			name:     "regular list, regular elem",
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			sls:      [2]int{3, 7},
			expected: []string{"c", "d", "e", "f", "g"},
			err:      nil,
		},
		{
			name:     "regular list, regular elem",
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			sls:      [2]int{3, 10},
			expected: []string{"c", "d", "e", "f", "g", "h", "i", "j"},
			err:      nil,
		},
		{
			name:     "regular list, elems are regular and equal",
			input:    []string{"A", "B", "C"},
			sls:      [2]int{2, 2},
			expected: []string{"B"},
			err:      nil,
		},
		{
			name:     "regular list, elems are out of range",
			input:    []string{"L", "M", "N", "O", "P", "Q"},
			sls:      [2]int{0, 5}, // should start since 1
			expected: []string{},
			err:      ErrOutOfRange,
		},
		{
			name:     "regular list, elems are out of range",
			input:    []string{"L", "M", "N", "O", "P", "Q"},
			sls:      [2]int{3, 7}, // should start at 6
			expected: []string{},
			err:      ErrOutOfRange,
		},
		{
			name:     "regular list, 1th elem is greater than 2th",
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"},
			sls:      [2]int{7, 3},
			expected: []string{},
			err:      ErrIncorrectSls,
		},
		{
			name:     "empty list, regular elems",
			input:    []string{},
			sls:      [2]int{3, 7},
			expected: []string{},
			err:      ErrInputListIsEmpty,
		},
		{
			name:     "regular list, elems are negative",
			input:    []string{"A", "B", "C", "D", "F"},
			sls:      [2]int{-2, -3},
			expected: []string{},
			err:      ErrOutOfRange,
		},
	}

	for _, tstCase := range tstCases {
		t.Run(tstCase.name, func(t *testing.T) {
			got, err := ExtractSlice(tstCase.input, tstCase.sls)

			if !reflect.DeepEqual(got, tstCase.expected) {
				t.Errorf("ExtractSlice(): got %v, want %v", got, tstCase.expected)
			}

			if err != tstCase.err {
				t.Errorf("ExtractSlice(): got %v, want %v", err, tstCase.err)
			}
		})
	}
}
