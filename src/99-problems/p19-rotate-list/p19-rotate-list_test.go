package rotatelist

import (
	"reflect"
	"testing"
)

type TstCase struct {
	name     string
	input    []string
	nPlaces  int
	expected []string
	err      error
}

func TestRotateList(t *testing.T) {
	tstCases := []TstCase{
		{
			name:     "empty list, N is out of range",
			input:    []string{},
			nPlaces:  10,
			expected: []string{},
			err:      ErrListIsEmpty,
		},
		{
			name:     "regular list, N == 0",
			input:    []string{"A", "B", "C"},
			nPlaces:  0,
			expected: []string{"A", "B", "C"},
			err:      nil,
		},
		{
			name:     "regular list, regular N == 1",
			input:    []string{"a", "b", "c"},
			nPlaces:  1,
			expected: []string{"b", "c", "a"},
			err:      nil,
		},
		{
			name:     "regular list, regular N == 3",
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h"},
			nPlaces:  3,
			expected: []string{"d", "e", "f", "g", "h", "a", "b", "c"},
			err:      nil,
		},
		{
			name:     "regular list, N == len(list)",
			input:    []string{"a", "b", "c", "d"},
			nPlaces:  4,
			expected: []string{"a", "b", "c", "d"},
			err:      nil,
		},
		{
			name:     "regular list, negative N",
			input:    []string{"a", "b", "c"},
			nPlaces:  -1,
			expected: []string{},
			err:      ErrInvalidN,
		},
		{
			name:     "regular list, N > len(list)",
			input:    []string{"a", "b", "c"},
			nPlaces:  100,
			expected: []string{},
			err:      ErrInvalidN,
		},
		{
			name:     "regular list, regular N == 15",
			input:    []string{"a", "b", "c", "d", "e", "f", "g", "h"},
			nPlaces:  15,
			expected: []string{"h", "a", "b", "c", "d", "e", "f", "g"},
			err:      nil,
		},
	}

	for _, tst := range tstCases {
		t.Run(tst.name, func(t *testing.T) {
			got, err := RotateList(tst.input, tst.nPlaces)

			if !reflect.DeepEqual(got, tst.expected) {
				t.Errorf("got %v, want %v", got, tst.expected)
			}

			if err != tst.err {
				t.Errorf("got %v, want %v", err, tst.err)
			}
		})
	}
}
