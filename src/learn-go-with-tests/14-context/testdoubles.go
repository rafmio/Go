package ctxt

import (
	"testing"
	"time"
)

// SpyStore allows you to simulate a store and see how its used
type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

// Fetch returns response after a short delay
func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}
