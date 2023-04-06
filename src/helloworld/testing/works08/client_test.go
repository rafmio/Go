package client

import (
	"fmt"
	"testing"
	"time"
)

func TestRunClient(t *testing.T) {
	currentDate := time.Now()
	want := currentDate.Format("02-01-2006")

	go runServer()

	got := runClient()
	if got != want {
		t.Errorf("incorrect return: wanted: %q, got: %q\n", want, got)
	} else {
		fmt.Println("Everything is OK")
		fmt.Printf("Wanted: %q, got: %q\n", want, got)
	}
}

func BenchmarkRunClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = runClient()
	}
}

// go test
// go test -v
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -bench=.
// go test -bench=generateAcc
