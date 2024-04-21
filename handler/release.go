package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"sync"
	"time"

	"github.com/docker/go-units"
	"github.com/mmcdole/gofeed"
	"github.com/recoilme/slowpoke"

	"rss.knaben.eu/entity"
	"rss.knaben.eu/parser"
	"rss.knaben.eu/service"
)

const db = "db/knaben.db"

var rLink = regexp.MustCompile(`<a href="(.*?)">`)
var rSize = regexp.MustCompile(`Size: (.*?)<br />`)

var api = service.NewIMDB()

func Parse(items []*gofeed.Item) {
	limiter := New(10)

	var lk sync.Mutex

	for _, item := range items {
		// fmt.Println(item.Title)

		limiter.Add()

		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)

		go func(item *gofeed.Item) {
			defer limiter.Done()

			parser := parser.NewParser()
			release := parser.Parse(item.Title)

			Enrich(release, item)

			lk.Lock()
			defer lk.Unlock()

			fmt.Println(release)
			// fmt.Println(release.ShowRSS())
		}(item)

		// if i > 10 {
		// break
		// }
	}

	limiter.Wait()
}

func Enrich(release *entity.Release, item *gofeed.Item) {
	if exist, err := slowpoke.Has(db, []byte(item.Title)); err != nil {
		panic(err)
	} else if exist {
		if v, err := slowpoke.Get(db, []byte(item.Title)); err != nil {
			panic(err)
		} else {
			if err := json.Unmarshal(v, release); err != nil {
				panic(err)
			}
		}
	} else {
		query := release.Title

		if release.Year > 0 {
			query = fmt.Sprintf("%s %d", release.Title, release.Year)
		}

		if res := api.Suggestion(query); len(res.Data) > 0 {
			for _, data := range res.Data {
				if data.Type != entity.Movie && data.Type != entity.TVSeries && data.Type != entity.TVMiniSeries {
					break
				}

				release.IMDB.ID = data.ID

				release.IMDB.Title = data.Name
				release.IMDB.Year = data.Year

				release.IMDB.Type = data.Type

				switch data.Type {
				case entity.Movie:
					{
						if res := api.Title(release.IMDB.ID); res != nil {
							data := res.Props.PageProps.AboveTheFoldData

							releaseDate := time.Date(data.ReleaseDate.Year, time.Month(data.ReleaseDate.Month), data.ReleaseDate.Day, 0, 0, 0, 0, time.UTC)

							release.IMDB.ID = data.ID
							release.IMDB.Type = data.TitleType.ID

							release.IMDB.Title = data.TitleText.Text
							release.IMDB.Year = data.ReleaseYear.Year

							release.IMDB.ReleaseDate = &releaseDate

							release.IMDB.Runtime = data.Runtime.Seconds

							release.IMDB.Rating = data.Certificate.Rating

							for _, v := range data.TitleGenres.Genres {
								release.IMDB.Genres = append(release.IMDB.Genres, v.Genre.Text)
							}

							release.IMDB.Plot = data.Plot.PlotText.PlainText

							release.IMDB.Ratings = data.RatingsSummary.AggregateRating
							release.IMDB.Metascore = data.Metacritic.Metascore.Score

							for _, v := range data.CountriesOfOrigin.Countries {
								release.IMDB.Countries = append(release.IMDB.Countries, v.ID)
							}
						}
					}
				case entity.TVSeries,
					entity.TVMiniSeries:
					{
						var season int
						var episode int

						fmt.Sscanf(release.Season, "S%d", &season)
						fmt.Sscanf(release.Episode, "E%d", &episode)

						if res := api.Episodes(release.IMDB.ID, season); res != nil {
							data := res.Props.PageProps.ContentData.EntityMetadata

							release.IMDB.ID = data.ID
							release.IMDB.Type = data.TitleType.ID

							release.IMDB.Title = data.TitleText.Text
							release.IMDB.Year = data.ReleaseYear.Year

							release.IMDB.Runtime = data.Runtime.Seconds

							release.IMDB.Rating = data.Certificate.Rating

							for _, v := range data.TitleGenres.Genres {
								release.IMDB.Genres = append(release.IMDB.Genres, v.Genre.Text)
							}

							release.IMDB.Plot = data.Plot.PlotText.PlainText

							release.IMDB.Ratings = data.RatingsSummary.AggregateRating

							for _, v := range data.CountriesOfOrigin.Countries {
								release.IMDB.Countries = append(release.IMDB.Countries, v.ID)
							}

							if episode > 0 && res.Props.PageProps.Section.Episodes.Total >= episode {
								data := res.Props.PageProps.Section.Episodes.Items[episode]

								releaseDate := time.Date(data.ReleaseDate.Year, time.Month(data.ReleaseDate.Month), data.ReleaseDate.Day, 0, 0, 0, 0, time.UTC)

								release.IMDB.ID = data.ID
								release.IMDB.Type = data.Type

								release.IMDB.Season = data.Season
								release.IMDB.Episode = data.Episode

								release.IMDB.Title = data.TitleText
								release.IMDB.Year = data.ReleaseYear

								release.IMDB.ReleaseDate = &releaseDate

								release.IMDB.Plot = data.Plot

								release.IMDB.Ratings = data.AggregateRating
							}
						}
					}
				}

				break
			}
		}

		release.RSS.GUID = item.GUID

		release.RSS.Title = item.Title
		release.RSS.Link = item.Link

		release.RSS.Categories = item.Categories
		release.RSS.Source = item.Author.Name

		if date, err := time.Parse(time.RFC1123Z, item.Published); err != nil {
			panic(err)
		} else {
			utc := date.UTC()
			release.RSS.Published = &utc // .Format(time.RFC3339)
		}

		description := item.Description
		link := rLink.FindStringSubmatch(description)[1]
		release.Link = link

		size := rSize.FindStringSubmatch(description)[1]
		if size, err := units.FromHumanSize(size); err != nil {
			panic(err)
		} else {
			release.Size = size
		}

		if v, err := json.Marshal(release); err != nil {
			panic(err)
		} else {
			slowpoke.Set(db, []byte(item.Title), v)
		}
	}
}
