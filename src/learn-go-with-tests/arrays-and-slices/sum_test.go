package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

// $ go test
// $ go test -v
// $ go test -cover

func ExampleSum() {
	sizesArray := []int{16, 32, 64, 128, 256, 512, 1024}
	sum := Sum(sizesArray)
	fmt.Println(sum)
}

// go test -v

func BenchmarkSum(b *testing.B) {
	numArray := []int{16, 32, 64, 128, 256, 512, 1024}
	_ = Sum(numArray)
	for i := 0; i < b.N; i++ {
		Sum(numArray)
	}

}

// $ go test -bench=.

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9}, []int{100, 200, 300})
	want := []int{3, 9, 600}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{200, 300})
		want := []int{2, 9, 300}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
