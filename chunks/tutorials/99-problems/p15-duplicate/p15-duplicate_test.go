package duplicate

import (
	"reflect"
	"testing"
)

var inputs map[int][]string = map[int][]string{
	0: []string{"a"},
	1: []string{"a", "b", "c", "d", "e", "f"},
	2: []string{"a", "b", "c", "d", "e"},
	3: []string{"a", "b", "c", "d"},
	4: []string{"a", "b", "c"},
	5: []string{"a", "b"},
}

var outputs map[int][]string = map[int][]string{
	0: []string{},
	1: []string{"a", "b", "c", "d", "e", "f"},
	2: []string{"a", "b", "c", "d", "e", "a", "b", "c", "d", "e"},
	3: []string{"a", "b", "c", "d", "a", "b", "c", "d", "a", "b", "c", "d"},
	4: []string{"a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c"},
	5: []string{"a", "b", "a", "b", "a", "b", "a", "b", "a", "b"},
}

func TestDuplicate(t *testing.T) {
	for i, val := range inputs {
		t.Run("pass values:", func(t *testing.T) {
			got := Duplicate(i, val)
			want := outputs[i]
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}
