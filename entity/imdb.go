package entity

import (
	"strings"
	"time"
)

const IMDB_URI = "https://www.imdb.com/title/"

type Type string

const (
	Movie        Type = "movie"
	TVSeries     Type = "tvSeries"
	TVMiniSeries Type = "tvMiniSeries"
)

type IMDB struct {
	ID   string `json:"id,omitempty"`
	Type Type   `json:"type,omitempty"`

	Season  string `json:"season,omitempty"`
	Episode string `json:"episode,omitempty"`

	Title string `json:"titleText,omitempty"`
	Year  int    `json:"releaseYear,omitempty"`

	ReleaseDate *time.Time `json:"releaseDate,omitempty"`

	Runtime int `json:"runtime,omitempty"` // 5580 => "1h 33m"

	Rating string `json:"rating,omitempty"` // "R"

	Genres []string `json:"genres,omitempty"` // ["Horror"]

	Plot string `json:"plot,omitempty"` // "A live television broadcast in 1977 goes horribly wrong, unleashing evil into the nation's living rooms."

	Ratings   float64 `json:"ratings,omitempty"`   // 7.3
	Metascore int     `json:"metascore,omitempty"` // 72
	// > 60 (GREEN), < 40 (RED), >= 40 && <= 60 ORANGE

	Countries []string `json:"countries,omitempty"` // ["AU"]
}

func (i *IMDB) URI() string {
	return IMDB_URI + i.ID + "/"
}

func (i *IMDB) Duration() (duration string) {
	if i.Runtime > 0 {
		d := time.Duration(i.Runtime) * time.Second

		duration = i.shortDuration(d)
	}

	return
}

func (i *IMDB) shortDuration(d time.Duration) (s string) {
	s = d.String()

	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}

	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}

	return
}
