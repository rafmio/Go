package insertelem

import (
	"reflect"
	"testing"
)

type TstCase struct {
	name     string
	inputSls []string
	position int
	element  string
	expected []string
}

func TestInsertElement(t *testing.T) {
	tstCases := []TstCase{
		{
			name:     "inputSls 4 elem, position 3, element 'E'",
			inputSls: []string{"A", "B", "C", "D"},
			position: 3,
			element:  "E",
			expected: []string{"A", "B", "E", "C", "D"},
		},
		{
			name:     "inputSls 4 elem, position 0, element 'E'",
			inputSls: []string{"A", "B", "C", "D"},
			position: 0,
			element:  "E",
			expected: []string{"E", "A", "B", "C", "D"},
		},
		{
			name:     "inputSls 4 elem, position len(inputSls), elem'E'",
			inputSls: []string{"A", "B", "C", "D"},
			position: 5,
			element:  "E",
			expected: []string{"A", "B", "C", "D", "E"},
		},
		{
			name:     "inputSls 6 elem, position 4, insert 'E'",
			inputSls: []string{"a", "b", "c", "d", "e", "f"},
			position: 4,
			element:  "E",
			expected: []string{"a", "b", "c", "E", "d", "e", "f"},
		},
	}

	for _, tstCase := range tstCases {
		t.Run(tstCase.name, func(t *testing.T) {
			got := InsertElement(tstCase.inputSls, tstCase.position, tstCase.element)

			if !reflect.DeepEqual(got, tstCase.expected) {
				t.Errorf("InsertElement() = %v, want %v", got, tstCase.expected)
			}
		})
	}
}
