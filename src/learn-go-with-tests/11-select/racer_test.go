package racer

import (
	"testing"
)

func TestRacer(t *testing.T) {
	slowURL := "https://www.google.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select
