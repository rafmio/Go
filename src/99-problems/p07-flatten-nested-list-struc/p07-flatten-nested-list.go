package flattenNL

import (
	"testing"
)

func TestFlattenNL(t *testing.T) {
	tstCases := []struct {
		name     string
		input    interface{}
		expected []interface{}
	}{
		{"shouldFlattenAListOfList", []string{"a", []string{"b", []string{"c", "d"}}, "e"}, []interface{}{"a", "b", "c", "d", "e"}},
		{"shouldFlattenDeepNestedLists", []string{"a", []string{"b", []string{"c", []string{"d", "e", []string{"f", "g"}}}, "h"}}, []interface{}{"a", "b", "c", "d", "e", "f", "g", "h"}},
		{"shouldReturnEmptyListWhenTryingToFlattenAnEmptyList", []string{}, []interface{}{}},
	}

}
