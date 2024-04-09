package packdup

import (
	"reflect"
	"testing"
)

type DupStr struct {
	origin []string
	packed [][]string
}

func TestPackDupStr(t *testing.T) {
	strCases := []DupStr{
		{
			origin: []string{"a", "a", "a", "a", "b", "c", "c", "d", "e", "e", "e", "e"},
			packed: [][]string{[]string{"a", "a", "a", "a"}, []string{"b"}, []string{"c", "c"}, []string{"d"}, []string{"e", "e", "e", "e"}},
		},
		{
			origin: []string{"f", "f", "f", "g", "h", "i", "i", "j", "k", "l", "m", "m"},
			packed: [][]string{[]string{"f", "f", "f"}, []string{"g"}, []string{"h"}, []string{"i", "i"}, []string{"j"}, []string{"k"}, []string{"l"}, []string{"m", "m"}},
		},
		{
			origin: []string{"A", "A", "A", "B", "C", "C", "D", "E", "F", "F", "F", "F"},
			packed: [][]string{[]string{"A", "A", "A"}, []string{"B"}, []string{"C", "C"}, []string{"D"}, []string{"E"}, []string{"F", "F", "F", "F"}},
		},
	}

	t.Run("passing strings", func(t *testing.T) {
		for _, strCase := range strCases {
			got := PackDupStr(strCase.origin)
			want := strCase.packed

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		}
	})
}
