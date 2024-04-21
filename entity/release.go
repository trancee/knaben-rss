package entity

import (
	"encoding/json"
	"fmt"

	"github.com/docker/go-units"
)

// Feature.Title.<YEAR>.<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP
// Weekly.TV.Show.[COUNTRY_CODE].[YEAR].SXXEXX[Episode.Part].[Episode.Title].<TAGS>.[LANGUAGE].<RESOLUTION>.<FORMAT>.<x264|x265>-GROUP

type Release struct {
	Title   string `json:"title"`
	Country string `json:"country,omitempty"`
	Year    int64  `json:"year,omitempty"`

	Season       string `json:"season,omitempty"`
	Episode      string `json:"episode,omitempty"`
	EpisodeTitle string `json:"episodeTitle,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Language []string `json:"language,omitempty"`

	Resolution []string `json:"resolution,omitempty"`
	Quality    []string `json:"quality,omitempty"`

	Network []string `json:"network,omitempty"`

	Source []string `json:"source,omitempty"`

	Audio    []string `json:"audio,omitempty"`
	Channels []string `json:"channels,omitempty"`

	Compression []string `json:"compression,omitempty"`

	Group []string `json:"group,omitempty"`

	Size int64  `json:"size,omitempty"`
	Link string `json:"link,omitempty"`

	RSS  RSS  `json:"rss,omitempty"`
	IMDB IMDB `json:"imdb,omitempty"`
}

func (r *Release) GetResolution() string {
	var s string
	for _, v := range r.Resolution {
		if s != v {
			s = v
		}
	}
	return s
}

func (r *Release) GetSource() string {
	var s string
	for _, v := range r.Source {
		if s != v {
			s = v
		}
	}
	return s
}

func (r *Release) GetChannels() string {
	var s string
	for _, v := range r.Channels {
		if s != v {
			s = v
		}
	}
	return s
}

func (r *Release) GetCompression() string {
	var s string
	for _, v := range r.Compression {
		if s != v && v != "HEVC" {
			s = v
		}
	}
	return s
}

func (r *Release) GetSize() string {
	return units.HumanSize(float64(r.Size))
}

func (r *Release) String() string {
	title := r.Title

	if r.Year > 0 {
		title = fmt.Sprintf("%s (%d)", title, r.Year)
	}

	show := ""

	if len(r.Season) > 0 {
		show = fmt.Sprintf("%3s%3s", r.Season, r.Episode)
	}

	metadata := fmt.Sprintf(
		"[%5s] [%-6s] [%3s] [%4s] %8s",
		r.GetResolution(),

		r.GetSource(),

		r.GetChannels(),

		r.GetCompression(),

		r.GetSize(),
	)

	ratings := "　"
	duration := "　"

	link := r.RSS.Link

	addendum := ""

	if len(r.IMDB.ID) > 0 {
		title = r.IMDB.Title

		link = r.IMDB.URI()

		addendum = fmt.Sprintf("\n%-125s %s", r.RSS.Title, r.RSS.Link)
	}

	switch r.IMDB.Type {
	case Movie:
		{
			title = r.IMDB.Title

			if r.IMDB.Year > 0 {
				title = fmt.Sprintf("%s (%d)", title, r.IMDB.Year)
			}
		}
	case TVSeries,
		TVMiniSeries:
		{
			season := r.Season
			if v := r.IMDB.Season; len(v) > 0 {
				season = fmt.Sprintf("S%02s", v)
			}
			episode := r.Episode
			if v := r.IMDB.Episode; len(v) > 0 {
				episode = fmt.Sprintf("E%02s", v)
			}

			if len(season) > 0 || len(episode) > 0 {
				show = fmt.Sprintf("%3s%3s", season, episode)
			}
		}
	}

	if r.IMDB.Ratings > 0 {
		ratings = fmt.Sprintf("⭐ %.1f", r.IMDB.Ratings)
	}
	if r.IMDB.Runtime > 0 {
		duration = fmt.Sprintf("🕑 %6s", r.IMDB.Duration())
	}

	return fmt.Sprintf(
		"%-60s%6s %s %-5s %-8s 🔗 %s%s",

		title,
		show,

		metadata,

		ratings,
		duration,

		link,

		addendum,
	)
}

func (r *Release) ShowRSS() string {
	return fmt.Sprintf(
		"%-74s %s %s %s",

		r.RSS.Title,

		r.RSS.Categories,
		r.RSS.Source,

		r.RSS.Link,
	)
}

func (r *Release) ShowIMDB() string {
	return fmt.Sprintf(
		"%-74s %s %.1f %s %s %s",

		fmt.Sprintf("%s (%d)", r.IMDB.Title, r.IMDB.Year),

		r.IMDB.Genres,

		r.IMDB.Ratings,
		r.IMDB.Rating,

		r.IMDB.Duration(),

		r.IMDB.URI(),
	)
}

func (r *Release) JSON() string {
	result, _ := json.Marshal(r)
	return string(result)
}

func (r *Release) Render() string {
	return Render(r)
}
