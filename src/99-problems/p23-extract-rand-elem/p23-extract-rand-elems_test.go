package extractrand

import (
	"testing"
)

type TstCase struct {
	name  string
	input []string
	num   int
	err   error
}

func TestExtractRandElems(t *testing.T) {
	tstCases := []TstCase{
		{
			name:  "10 elems, 5 nums",
			input: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			num:   5,
			err:   nil,
		},
		{
			name:  "10 elems, 10 nums",
			input: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			num:   10,
			err:   nil,
		},
		{
			name:  "5 elems, 1 num",
			input: []string{"a", "b", "c", "d", "e"},
			num:   1,
			err:   nil,
		},
		{
			name:  "5 elems, 0 nums",
			input: []string{"a", "b", "c", "d", "e"},
			num:   0,
			err:   nil,
		},
		{
			name:  "empty input",
			input: []string{},
			num:   3,
			err:   ErrEmptyInputList,
		},
		{
			name:  "3 elems, 10 nums (out of range)",
			input: []string{"a", "b", "c"},
			num:   10,
			err:   ErrIncorrectNum,
		},
		{
			name:  "3 elems, num -1 (negative num)",
			input: []string{"a", "b", "c"},
			num:   -1,
			err:   ErrIncorrectNum,
		},
	}

	for _, tstCase := range tstCases {

		if tstCase.num <= 0 || tstCase.num > len(tstCase.input) {
			continue
		}

		t.Run(tstCase.name, func(t *testing.T) {
			got, err := ExtractRandElems(tstCase.input, tstCase.num)

			if len(got) != tstCase.num {
				t.Errorf("ExtractRandElems(): incorrect number of elements: got = %d, want %d", len(got), len(tstCase.input))
			}
			if err != tstCase.err {
				t.Errorf("ExtractRandElems(): error = %v, want %v", err, tstCase.err)
			}
		})

	}
}
