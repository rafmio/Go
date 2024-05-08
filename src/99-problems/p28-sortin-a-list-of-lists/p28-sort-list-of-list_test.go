package sortlistoflist

import (
	"reflect"
	"testing"
)

type TstCase struct {
	Input    [][]string
	Expected [][]string
	err      error
}

func TestSortListOfList(t *testing.T) {
	tstCases := []TstCase{
		TstCase{
			Input: [][]string{
				{"a", "b", "c"},
				{"d", "e"},
				{"f"},
			},
			Expected: [][]string{
				{"f"},
				{"d", "e"},
				{"a", "b", "c"},
			},
			err: nil,
		},
		TstCase{
			Input: [][]string{
				{"F", "F"},
				{"F", "F", "F", "F"},
				{"F"},
				{"F", "F", "F"},
			},
			Expected: [][]string{
				{"F"},
				{"F", "F"},
				{"F", "F", "F"},
				{"F", "F", "F", "F"},
			},
			err: nil,
		},
		TstCase{
			Input: [][]string{
				{"1", "2", "3", "4", "5", "6"},
				{"1", "2", "3", "4"},
				{"1", "2", "3", "4", "5"},
			},
			Expected: [][]string{
				{"1", "2", "3", "4"},
				{"1", "2", "3", "4", "5"},
				{"1", "2", "3", "4", "5", "6"},
			},
			err: nil,
		},
		TstCase{
			Input: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
				{"1", "2", "3", "4"},
			},
			Expected: [][]string{
				{"1", "2"},
				{"1", "2", "3"},
				{"1", "2", "3", "4"},
			},
			err: nil,
		},
		TstCase{
			Input:    [][]string{},
			Expected: [][]string{},
			err:      ErrEmptyList,
		},
	}

	for _, tstCase := range tstCases {
		t.Run("run SortListOfList", func(t *testing.T) {
			result, err := SortListOfList(tstCase.Input)
			if err != tstCase.err {
				t.Errorf("expected error %v, got %v", tstCase.err, err)
			}
			if !reflect.DeepEqual(result, tstCase.Expected) {
				t.Errorf("expected %v, got %v", tstCase.Expected, result)
			}
		})
	}
}
