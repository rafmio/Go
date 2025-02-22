// https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package
package math

import (
	"fmt"
	"testing"
)

// func TestAdd(t *testing.T) {
// 	got := Add(4, 6)
// 	want := 10

// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
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

/*
Run test:
$ go test -v

Calculate coverage:
$ go test -coverprofile=coverage.out

To present the result in a web-browser:
$ go tool cover -html=coverage.out
*/

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

// Run benchmark test:
// $ go test -bench=.

// We can declare benchmark functions explictly:
// go test -bench=Add

func ExampleAdd() {
	fmt.Println(Add(400, 600))
}
