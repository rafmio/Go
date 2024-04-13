package modEnc

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
				"b": 1,
				"d": 1,
			},
		},
		{
			origin: []string{"f", "f", "f", "g", "h", "i", "i", "j", "k", "l", "m", "m"},
			encodedMap: map[string]int{
				"g": 1,
				"h": 1,
				"j": 1,
				"k": 1,
				"l": 1,
			},
		},
		{
			origin: []string{"A", "A", "A", "B", "C", "C", "D", "E", "F", "F", "F", "F"},
			encodedMap: map[string]int{
				"B": 1,
				"D": 1,
				"E": 1,
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
