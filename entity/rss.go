package entity

import "time"

type RSS struct {
	GUID string `json:"guid,omitempty"`

	Title string `json:"title,omitempty"`
	Link  string `json:"link,omitempty"`

	Categories []string `json:"categories,omitempty"`
	Source     string   `json:"source,omitempty"`

	Published *time.Time `json:"published,omitempty"`
}
