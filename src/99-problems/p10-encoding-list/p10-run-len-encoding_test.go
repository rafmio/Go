// origin: []string{"a", "a", "a", "a", "b", "c", "c", "d", "e", "e", "e", "e"},
// packed: [][]string{[]string{"a", "a", "a", "a"}, []string{"b"}, []string{"c", "c"}, []string{"d"}, []string{"e", "e", "e", "e"}},
// origin: []string{"f", "f", "f", "g", "h", "i", "i", "j", "k", "l", "m", "m"},
// packed: [][]string{[]string{"f", "f", "f"}, []string{"g"}, []string{"h"}, []string{"i", "i"}, []string{"j"}, []string{"k"}, []string{"l"}, []string{"m", "m"}},
// origin: []string{"A", "A", "A", "B", "C", "C", "D", "E", "F", "F", "F", "F"},
// packed: [][]string{[]string{"A", "A", "A"}, []string{"B"}, []string{"C", "C"}, []string{"D"}, []string{"E"}, []string{"F", "F", "F", "F"}},
package runLenEnc

import (
	"reflect"
	"testing"
)

type EncodedRes struct {
	origin     []string
	encodedMap map[string]int
}

func TestRunLengthEncoding(t *testing.T) {
	strCases := []EncodedRes{
		{
			origin: []string{"a", "a", "a", "a", "b", "c", "c", "d", "e", "e", "e", "e"},
			encodedMap: map[string]int{
				"a": 4,
				"b": 1,
				"c": 2,
				"d": 1,
				"e": 4,
			},
		},
		{
			origin: []string{"f", "f", "f", "g", "h", "i", "i", "j", "k", "l", "m", "m"},
			encodedMap: map[string]int{
				"f": 3,
				"g": 1,
				"h": 1,
				"i": 2,
				"j": 1,
				"k": 1,
				"l": 1,
				"m": 2,
			},
		},
		{
			origin: []string{"A", "A", "A", "B", "C", "C", "D", "E", "F", "F", "F", "F"},
			encodedMap: map[string]int{
				"A": 3,
				"B": 1,
				"C": 2,
				"D": 1,
				"E": 1,
				"F": 4,
			},
		},
	}

	for _, strCase := range strCases {
		t.Run("run every case", func(t *testing.T) {
			got := RunLengthEncoding(strCase.origin)
			want := strCase.encodedMap

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
