// https://github.com/quii/learn-go-with-tests/blob/main/context/v2/context.go
package ctxt

import "net/http"

// Store fetches data
type Store interface {
	Fetch() string
	Cancel()
}

// Server returns a handler for calling Store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // http.Context() function, return context.Context

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Frpint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
