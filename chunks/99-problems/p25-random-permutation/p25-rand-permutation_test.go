package randperm

import (
	"fmt"
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
			origin := make([]string, len(tstCase))
			copy(origin, tstCase)

			got := RandPerm(tstCase)

			if len(got) != len(tstCase) {
				t.Errorf("incorrect length")
				fmt.Println(got)
			}

			if reflect.DeepEqual(got, origin) {
				t.Errorf("the itemw in the list has not been permutated")
			}
		})
	}
}
