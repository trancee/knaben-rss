package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/recoilme/slowpoke"

	"rss.knaben.eu/entity"
)

const IMDB_API_URI = "https://v3.sg.media-imdb.com/suggestion/x/"

// const IMDB_URI = "https://www.imdb.com/title/"

type IIMDB interface {
	Suggestion(query string) *SuggestionResponse

	Title(id string) *AboveTheFoldData
	Episodes(id string, season int) *ContentData

	URI(id string) string
}

type IMDB struct {
	dbSuggestions string

	dbTitles   string
	dbEpisodes string
}

func NewIMDB() IIMDB {
	return &IMDB{
		dbSuggestions: "db/imdb.suggestions.db",

		dbTitles:   "db/imdb.titles.db",
		dbEpisodes: "db/imdb.episodes.db",
	}
}

func (i *IMDB) Suggestion(query string) *SuggestionResponse {
	res := &SuggestionResponse{}

	query = strings.TrimSpace(strings.ToLower(query))

	_id := Hash(s2b(query))

	if exists, err := slowpoke.Has(i.dbSuggestions, _id); err != nil {
		panic(err)
	} else if exists {
		if body, err := slowpoke.Get(i.dbSuggestions, _id); err != nil {
			panic(err)
		} else {
			if err := json.Unmarshal(body, res); err != nil {
				panic(err)
			}

			if len(res.Data) > 0 {
				return res
			}
		}
	}

	if req, err := http.NewRequest("GET", IMDB_API_URI+query+".json", nil); err != nil {
		panic(err)
	} else {
		client := &http.Client{}

		if resp, err := client.Do(req); err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()

			if body, err := io.ReadAll(resp.Body); err != nil {
				panic(err)
			} else {
				if err := json.Unmarshal(body, res); err != nil {
					panic(err)
				}

				// fmt.Println(query, b2s(body))

				if err := slowpoke.Set(i.dbSuggestions, _id, body); err != nil {
					panic(err)
				}
			}
		}
	}

	return res
}

func (i *IMDB) Title(id string) *AboveTheFoldData {
	res := &AboveTheFoldData{}

	_id := Hash(s2b(id))

	if exists, err := slowpoke.Has(i.dbTitles, _id); err != nil {
		panic(err)
	} else if exists {
		if body, err := slowpoke.Get(i.dbTitles, _id); err != nil {
			panic(err)
		} else {
			if err := json.Unmarshal(body, res); err != nil {
				panic(err)
			}

			// fmt.Println(id, b2s(body))

			return res
		}
	}

	if req, err := http.NewRequest("GET", entity.IMDB_URI+id, nil); err != nil {
		panic(err)
	} else {
		client := &http.Client{}

		req.Header.Add(
			"User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		)
		req.Header.Add(
			"Accept",
			"text/html",
		)
		req.Header.Add(
			"Accept-Language",
			"en-US",
		)

		if resp, err := client.Do(req); err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()

			if body, err := io.ReadAll(resp.Body); err != nil {
				panic(err)
			} else {
				doc := parse(b2s(body))

				if data := traverse(doc, "script", "id", "__NEXT_DATA__"); data != nil {
					if body, ok := text(data); ok {
						if err := json.Unmarshal(s2b(body), res); err != nil {
							panic(err)
						}

						// fmt.Println(id, b2s(body))

						if err := slowpoke.Set(i.dbTitles, _id, s2b(body)); err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}

	return res
}

func (i *IMDB) Episodes(id string, season int) *ContentData {
	res := &ContentData{}

	_id := Hash(s2b(id))

	if exists, err := slowpoke.Has(i.dbEpisodes, _id); err != nil {
		panic(err)
	} else if exists {
		if body, err := slowpoke.Get(i.dbEpisodes, _id); err != nil {
			panic(err)
		} else {
			if err := json.Unmarshal(body, res); err != nil {
				panic(err)
			}

			// fmt.Println(id, b2s(body))

			return res
		}
	}

	if req, err := http.NewRequest("GET", entity.IMDB_URI+id+"/"+"episodes", nil); err != nil {
		panic(err)
	} else {
		client := &http.Client{}

		req.Header.Add(
			"User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		)
		req.Header.Add(
			"Accept",
			"text/html",
		)
		req.Header.Add(
			"Accept-Language",
			"en-US",
		)

		if season > 0 {
			q := req.URL.Query()
			q.Add("season", fmt.Sprintf("%d", season))
			req.URL.RawQuery = q.Encode()
		}

		if resp, err := client.Do(req); err != nil {
			panic(err)
		} else {
			defer resp.Body.Close()

			if body, err := io.ReadAll(resp.Body); err != nil {
				panic(err)
			} else {
				doc := parse(b2s(body))

				if data := traverse(doc, "script", "id", "__NEXT_DATA__"); data != nil {
					if body, ok := text(data); ok {
						if err := json.Unmarshal(s2b(body), res); err != nil {
							panic(err)
						}

						// fmt.Println(id, b2s(body))

						if err := slowpoke.Set(i.dbEpisodes, _id, s2b(body)); err != nil {
							panic(err)
						}
					}
				}
			}
		}
	}

	return res
}

func (i *IMDB) URI(id string) string {
	return URI(id)
}

func URI(id string) string {
	return entity.IMDB_URI + id + "/"
}

type SuggestionResponse struct {
	Data []struct {
		ID string `json:"id"`

		Name string `json:"l"`
		Year int    `json:"y"`

		Type entity.Type `json:"qid"`
	} `json:"d"`
}

type AboveTheFoldData struct {
	Props struct {
		PageProps struct {
			AboveTheFoldData struct {
				ID string `json:"id"` // "tt14966898"

				TitleText struct {
					Text string `json:"text"` // "Late Night with the Devil"
				} `json:"titleText"`

				TitleType struct {
					ID   entity.Type `json:"id"`   // "movie"
					Text string      `json:"text"` // "Movie"

					IsSeries  bool `json:"isSeries"`  // false
					IsEpisode bool `json:"isEpisode"` // false
				} `json:"titleType"`

				Certificate struct {
					Rating string `json:"rating"` // "R"
				} `json:"certificate"`

				ReleaseYear struct {
					Year int `json:"year"` // 2023
				} `json:"releaseYear"`

				ReleaseDate struct {
					Month int `json:"month"` // 4
					Day   int `json:"day"`   // 14
					Year  int `json:"year"`  // 2024
				} `json:"releaseDate"`

				Runtime struct {
					Seconds int `json:"seconds"` // 5580

					DisplayableProperty struct {
						Value struct {
							PlainText string `json:"plainText"` // "1h 33m"
						} `json:"value"`
					} `json:"displayableProperty"`
				} `json:"runtime"`

				RatingsSummary struct {
					AggregateRating float64 `json:"aggregateRating"` // 7.3
					VoteCount       int     `json:"voteCount"`       // 13625
				} `json:"ratingsSummary"`

				Metacritic struct {
					Metascore struct {
						Score int `json:"score"` // 72
					} `json:"metascore"`
				} `json:"metacritic"`

				Genres struct {
					Genres []struct {
						ID   string `json:"id"`   // "Horror"
						Text string `json:"text"` // "Horror"
					} `json:"genres"`
				} `json:"genres"`

				TitleGenres struct {
					Genres []struct {
						Genre struct {
							Text string `json:"text"` // "Horror"
						} `json:"genre"`
					} `json:"genres"`
				} `json:"titleGenres"`

				Plot struct {
					PlotText struct {
						PlainText string `json:"plainText"` // "A live television broadcast in 1977 goes horribly wrong, unleashing evil into the nation's living rooms."
					} `json:"plotText"`
				} `json:"plot"`

				CountriesOfOrigin struct {
					Countries []struct {
						ID string `json:"id"` // "AU"
					} `json:"countries"`
				} `json:"countriesOfOrigin"`
			} `json:"aboveTheFoldData"`
		} `json:"pageProps"`
	} `json:"props"`
}

type ContentData struct {
	Props struct {
		PageProps struct {
			ContentData struct {
				EntityMetadata struct {
					ID string `json:"id"` // "tt15791752"

					Certificate struct {
						Rating string `json:"rating"` // "TV-14"
					} `json:"certificate"`

					Runtime struct {
						Seconds int `json:"seconds"` // 2460
					} `json:"runtime"`

					RatingsSummary struct {
						AggregateRating float64 `json:"aggregateRating"` // 5.7
					} `json:"ratingsSummary"`

					TitleGenres struct {
						Genres []struct {
							Genre struct {
								Text string `json:"text"` // "Animation"
							} `json:"genre"`
						} `json:"genres"`
					} `json:"titleGenres"`

					TitleType struct {
						ID   entity.Type `json:"id"`   // "tvSeries"
						Text string      `json:"text"` // "TV Series"

						IsSeries  bool `json:"isSeries"`  // true
						IsEpisode bool `json:"isEpisode"` // false
					} `json:"titleType"`

					TitleText struct {
						Text string `json:"text"` // "Grimsburg"
					} `json:"titleText"`

					ReleaseYear struct {
						Year int `json:"year"` // 2024
					} `json:"releaseYear"`

					Plot struct {
						PlotText struct {
							PlainText string `json:"plainText"` // "Marvin Flute, who might be the greatest detective ever, has one mystery he still can't crack: his family. He will follow every lead he's got to redeem himself with the ex-wife he never stopped loving."
						} `json:"plotText"`
					} `json:"plot"`

					CountriesOfOrigin struct {
						Countries []struct {
							ID string `json:"id"` // "US"
						} `json:"countries"`
					} `json:"countriesOfOrigin"`
				} `json:"entityMetadata"`
			} `json:"contentData"`

			Section struct {
				Episodes struct {
					Total int `json:"total"` // 13

					Items []struct {
						ID   string      `json:"id"`   // "tt31922591"
						Type entity.Type `json:"type"` // "tvEpisode"

						Season  string `json:"season"`  // "1"
						Episode string `json:"episode"` // "9"

						TitleText string `json:"titleText"` // "The Funaways"

						ReleaseDate struct {
							Month int `json:"month"` // 4
							Day   int `json:"day"`   // 14
							Year  int `json:"year"`  // 2024
						} `json:"releaseDate"`

						ReleaseYear int `json:"releaseYear"` // 2024

						Plot string `json:"plot"` // "Stan stages his disappearance with assistance from Dr. Pentos to stop Marvin from forgetting about him."

						AggregateRating float64 `json:"aggregateRating"` // 7.2
						VoteCount       int     `json:"voteCount"`       // 25
					} `json:"items"`
				} `json:"episodes"`
			} `json:"section"`
		} `json:"pageProps"`
	} `json:"props"`
}
