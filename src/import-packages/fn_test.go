package fn

import (
	"testing"
)

func TestFn(t *testing.T) {
	t.Run("fn", func(t *testing.T) {
		want := 10
		got := Fn(5)
		if got != want {
			t.Errorf("Fn() = %v, want %v", got, want)
		}
	})
}
