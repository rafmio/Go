// origin: []string{"a", "a", "a", "a", "b", "c", "c", "d", "e", "e", "e", "e"},
// origin: []string{"f", "f", "f", "g", "h", "i", "i", "j", "k", "l", "m", "m"},
// origin: []string{"A", "A", "A", "B", "C", "C", "D", "E", "F", "F", "F", "F"},
package count

import (
	"reflect"
	"testing"
)

type tstCase struct {
	origin []string
	result map[string]int
}

type GenTstCase[T comparable] struct {
	origin []T
	result map[T]int
}

func TestCount(t *testing.T) {
	tstCases := []tstCase{
		{
			origin: []string{"a", "a", "a", "a", "b", "c", "c", "d", "e", "e", "e", "e"},
			result: map[string]int{"a": 4, "b": 1, "c": 2, "d": 1, "e": 4},
		},
		{
			origin: []string{"f", "f", "f", "g", "h", "i", "i", "j", "k", "l", "m", "m"},
			result: map[string]int{"f": 3, "g": 1, "h": 1, "i": 2, "j": 1, "k": 1, "l": 1, "m": 2},
		},
		{
			origin: []string{"A", "A", "A", "B", "C", "C", "D", "E", "F", "F", "F", "F"},
			result: map[string]int{"A": 3, "B": 1, "C": 2, "D": 1, "E": 1, "F": 4},
		},
		{
			origin: []string{},
			result: map[string]int{},
		},
	}

	for _, tst := range tstCases {
		t.Run("pass test cases", func(t *testing.T) {
			got := Count(tst.origin)

			if !reflect.DeepEqual(got, tst.result) {
				t.Errorf("want: %v, got: %v", tst.result, got)
			}
		})
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count([]string{"A", "B", "b", "C", "T", "M", "T", "Z", "0", "I", "I"})
	}
}
