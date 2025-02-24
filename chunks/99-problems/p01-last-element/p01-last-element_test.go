package problems

import (
	"testing"
)

func TestFindLastElement(t *testing.T) {
	t.Run("finds the last element in the list", func(t *testing.T) {
		list := []string{"a", "b", "c", "d"}
		expected := "d"
		got := FindLastElement(list)

		if got != expected {
			t.Errorf("expected '%s' but got '%s'", expected, got)
		}
	})

	t.Run("returns 0 if the list is empty", func(t *testing.T) {
		list := []string{}
		expected := ""
		got := FindLastElement(list)

		if got != expected {
			t.Errorf("expected '%s' but got '%v'", expected, got)
		}
	})
}
