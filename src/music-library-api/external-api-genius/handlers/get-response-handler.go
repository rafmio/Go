package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"geniusapi/metadata"
	"geniusapi/models"
)

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// var queryParams models.QueryParams
	queryParams := new(models.QueryParams)
	err := json.NewDecoder(r.Body).Decode(queryParams)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	songDetails, err := metadata.GetSongMetadata(queryParams)
	if err != nil {
		log.Println("error getting song metadata")
		http.Error(w, "Error getting song metadata", http.StatusInternalServerError)
		return
	} else {
		log.Println("metadata for song has been extracted")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(songDetails)
}
