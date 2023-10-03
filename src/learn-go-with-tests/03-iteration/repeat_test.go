package iteration

import (
	"fmt"
	"os"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expecetd := "aaaaa"

	if repeated != expecetd {
		t.Errorf("expected %q but got %q", expecetd, repeated)
	} else {
		fmt.Fprintf(os.Stdout, "repeated: %s\n", repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("Y")
	}
}

// go test -bench=.

func ExampleRepeat() {
	res := Repeat("Hello-mello")
	fmt.Println(res)
}

// go test -v
