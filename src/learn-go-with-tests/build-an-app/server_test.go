package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlaers(t *testing.T) {
	t.Run("return Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		reponse := httptest.NewRecorder()

		PlayerServer(reponse, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
