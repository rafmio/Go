package swagger

import "net/http"

// curl -X POST \
//   http://localhost:8080/v1/songs \
//   -H 'Content-Type: application/json' \
//   -d '{
//   "group": "Muse",
//   "song": "Supermassive Black Hole"
// }'

func SongsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
