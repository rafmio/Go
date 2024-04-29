package randperm

import (
	"reflect"
	"testing"
)

func TestRandPerm(t *testing.T) {
	tstCases := [][]string{
		[]string{"a", "b", "c", "d", "e", "f"},
		[]string{"1", "2", "3", "4", "5", "6"},
		[]string{"Huggy", "Wuggy", "Kissy", "Missy"},
	}

	for _, tstCase := range tstCases {
		t.Run("run tstCase", func(t *testing.T) {
			got := RandPerm(tstCase)

			if reflect.DeepEqual(got, tstCase) || len(got) != len(tstCase) {
				t.Errorf("the item in the list has not been permutated or incorrect length")
			}
		})
	}
}
