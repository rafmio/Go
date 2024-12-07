package swagger

import (
	"log"
	"net/http"
)

// http://localhost:8080/v1/songs/search?title=Supermassive%20Black%20Hole&artist=Muse&release_date=10-05-2006
func SongsSearchGet(w http.ResponseWriter, r *http.Request) {
	log.Println("inside the SongsSearchGet() func") // DEBUG print

	// check if request's method is GET:
	if r.Method != "GET" {
		log.Println("the SongsSearchGet() receive wrong method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// parse query parameters:
	title := r.URL.Query().Get("title")
	artist := r.URL.Query().Get("artist")
	release_date := r.URL.Query().Get("release_date")

	// perform search using provided parameters:
	// ...

	// return search results as JSON:

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
