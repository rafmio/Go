package lastbutone

import (
	"testing"
)

func TestLastButOne(t *testing.T) {
	t.Run("send slice of ints", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got, _ := LastButOne([]interface{}{numbers})
		want := 4

		if got != want {
			t.Errorf("got %v, want %d", got, want)
		}
	})
}
