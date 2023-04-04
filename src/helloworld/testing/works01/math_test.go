package math

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

type subtractTest struct {
	arg1, arg2, expected int
}

var subtractTests = []subtractTest{
	subtractTest{10, 8, 2},
	subtractTest{100, 40, 60},
	subtractTest{1000000, 350000, 650000},
	subtractTest{750, 250, 500},
}

var addTests = []addTest{
	addTest{2, 3, 5},
	addTest{4, 8, 12},
	addTest{6, 9, 15},
	addTest{3, 10, 13},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func TestSubtract(t *testing.T) {
	for _, test := range subtractTests {
		if output := Subtract(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(340, 60)
	}
}

// go test
// go test -v
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -bench=.
// go test -bench=EstablishConnectionDB
