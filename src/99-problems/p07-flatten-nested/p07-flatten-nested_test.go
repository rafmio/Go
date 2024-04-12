package flatten

import (
	"reflect"
	"testing"
)

type StrCase struct {
	origin  [][]string
	flatted []string
}

func TestFlatten(t *testing.T) {
	strCases := []StrCase{
		{
			origin: [][]string{
				[]string{"a", "b", "c"},
				[]string{"d", "e", "f"},
			},
			flatted: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			origin: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g", "h", "i"},
			},
			flatted: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"},
		},
		{
			origin: [][]string{
				{"M", "a", "M", "a"},
				{"G", "r", "A", "n", "P", "a"},
				{"H"},
				{"o", "L", "a"},
			},
			flatted: []string{"M", "a", "M", "a", "G", "r", "A", "n", "P", "a", "H", "o", "L", "a"},
		},
		{
			origin: [][]string{
				{"H", "E", "R", "O"},
				{"o", "f"},
				{"t"},
				{"h", "e"},
				{"d", "a", "y"},
			},
			flatted: []string{"H", "E", "R", "O", "o", "f", "t", "h", "e", "d", "a", "y"},
		},
	}

	t.Run("runs test cases", func(t *testing.T) {
		for _, strCase := range strCases {
			flattened := Flatten(strCase.origin)
			if !reflect.DeepEqual(strCase.origin, strCase.flatted) {
				t.Errorf("got: %v, want: %v", flattened, strCase.flatted)
			}
		}
	})
}

// func TestFlatten(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		list [][]string
// 		want []string
// 	}{
// 		{
// 			name: "basic test",
// 			list: [][]string{
// 				[]string{"a", "b"},
// 				[]string{"c", "d"},
// 			},
// 			want: []string{"a", "b", "c", "d"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Flatten(tt.list); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Flatten() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
