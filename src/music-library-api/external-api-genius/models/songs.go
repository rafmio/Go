package models

type SongDetail struct {
	SeqNum      string // inner DB ID
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	// ReleaseDateForDisplay string `json:"release_date_for_display"`
	// ReleaseDate string `json:"releaseDate"`
	Artist string `json:"artist"`
	Text   string `json:"text,omitempty"`
	Lyrics string `json:"lyrics,omitempty"`
	Link   string `json:"link"`
}

type QueryParams struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
