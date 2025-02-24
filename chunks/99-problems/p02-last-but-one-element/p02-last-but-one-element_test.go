package lastbutone

import (
	"testing"
)

func TestLastButOne(t *testing.T) {
	t.Run("send slice of ints", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got, _ := LastButOne(numbers)
		want := 4

		if got != want {
			t.Errorf("got %v, want %d", got, want)
		}
	})

	t.Run("send empty slice of ints", func(t *testing.T) {
		numbers := []int{}
		want := 0

		got, err := LastButOne(numbers)
		if err != ErrNoSuchElementException { // There are no 'exceptions' in Go
			t.Errorf("got %v, want %v", err, ErrNoSuchElementException)
		}
		if got != 0 {
			t.Errorf("got %v, want %d", got, want)
		}
	})
}
