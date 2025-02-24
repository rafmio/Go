package main

import (
	"testing"
)

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result is incorrect, got; %s, want: %s\n", result, "Foo")
	}
}

func TestFooerTableDriven(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		// the table itself
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			_ = Fooer(j)
		}
	}
}

// go test
// go test -v
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -bench=.
// go test -bench=FuncName
