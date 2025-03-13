// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context
// https://github.com/quii/learn-go-with-tests/blob/main/context/v3/context.go
package ctxt

import (
	"context"
	"fmt"
	"net/http"
)

// Store fetches data
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server returns a handler for calling Store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
