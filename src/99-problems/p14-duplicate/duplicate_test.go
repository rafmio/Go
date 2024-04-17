package duplicate

import (
	"reflect"
	"testing"
)

type TstCases struct {
	input  []string
	output []string
}

func TestDuplicate(t *testing.T) {

	inputs := [][]string{
		[]string{"a", "b", "c", "d"},
		[]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
		[]string{"X", "Y", "Z"},
		[]string{},
	}

	outputs := [][]string{
		[]string{"a", "a", "b", "b", "c", "c", "d", "d"},
		[]string{"0", "0", "1", "1", "2", "2", "3", "3", "4", "4", "5", "5", "6", "6", "7", "7", "8", "8", "9", "9"},
		[]string{"X", "X", "Y", "Y", "Z", "Z"},
		[]string{},
	}

	tstCases := make([]TstCases, len(inputs))

	for i := 0; i < len(inputs); i++ {
		tstCases[i] = TstCases{
			input:  inputs[i],
			output: outputs[i],
		}
	}

	t.Run("pass cases", func(t *testing.T) {

		for _, tst := range tstCases {
			got := Duplicate(tst.input)

			if !reflect.DeepEqual(got, tst.output) {
				t.Errorf("got %v, want %v", got, tst.output)
			}
		}

	})
}
