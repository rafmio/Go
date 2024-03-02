package arrays

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of eny size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := make([]int, 50)
	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		numbers[i] = rand.Intn(1000)
	}

	for i := 0; i < b.N; i++ {
		_ = Sum(numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// go test -cover
// go test -bench=.
