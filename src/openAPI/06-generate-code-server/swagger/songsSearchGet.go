package swagger

import (
	"log"
	"net/http"

	"getcode/dbops"
	"getcode/models"
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

	queryParams := new(models.QueryParams)
	// parse query parameters:
	queryParams.Song = r.URL.Query().Get("title")
	queryParams.Group = r.URL.Query().Get("artist")
	queryParams.ReleaseDate = r.URL.Query().Get("release_date")

	// check if all fields are empty
	if queryParams.Song == "" && queryParams.Group == "" && queryParams.ReleaseDate == "" {
		log.Println("all parameters (fields) in request are empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Implement more complex validations like date format, etc.
	// TODO: Implement pagination and limit results to 10 per page.
	// TODO: Implement logging for API requests and responses.

	// search song
	log.Println("try to find song")
	songs, err := dbops.SongsSearchDB(queryParams)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
