package remove

import (
	"reflect"
	"testing"
)

type TstCase struct {
	name     string
	input    []string
	K        int
	expected []string
	err      error
}

func TestRemove(t *testing.T) {
	tstCases := []TstCase{
		{
			name:     "regular input, K = 0",
			input:    []string{"a", "b", "c", "d"},
			K:        0,
			expected: []string{"b", "c", "d"},
			err:      nil,
		},
		{
			name:     "regular input, K = 3",
			input:    []string{"A", "B", "C", "D", "E", "F"},
			K:        3,
			expected: []string{"A", "B", "C", "E", "F"},
			err:      nil,
		},
		{
			name:     "regular input, K = 5",
			input:    []string{"A", "B", "C", "D", "E", "F"},
			K:        5,
			expected: []string{"A", "B", "C", "D", "E"},
			err:      nil,
		},
		{
			name:     "regular input, K is negative",
			input:    []string{"A", "B", "C", "D", "E", "F"},
			K:        -1,
			expected: []string{},
			err:      ErrKIsIncorrect,
		},
		{
			name:     "regular input, K is out of range",
			input:    []string{"A", "B", "C", "D", "E", "F"},
			K:        10,
			expected: []string{},
			err:      ErrKIsIncorrect,
		},
		{
			name:     "empty input",
			input:    []string{},
			K:        3,
			expected: []string{},
			err:      ErrEmptyInput,
		},
	}

	for _, tstCase := range tstCases {
		t.Run(tstCase.name, func(t *testing.T) {
			got, err := Remove(tstCase.input, tstCase.K)
			if !reflect.DeepEqual(got, tstCase.expected) {
				t.Errorf("Remove() = %v, want %v", got, tstCase.expected)
			}

			if err != tstCase.err {
				t.Errorf("Remove() error = %v, want %v", err, tstCase.err)
			}
		})
	}
}
