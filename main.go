package main

import (
	"github.com/mmcdole/gofeed"
	"github.com/recoilme/slowpoke"

	"rss.knaben.eu/handler"
)

// const feedURL = "https://rss.knaben.eu/1080p x265/2001000//1337x|eztv|thepiratebay|yts|hidexxx"
// const feedURL = "https://rss.knaben.eu/x265/2001000|2003000//1337x|eztv|thepiratebay|yts|hidexxx"
const feedURL = "https://rss.knaben.eu/x265/3003000|3001000|3004000//1337x|eztv|thepiratebay|yts|hidexxx"

func main() {
	defer func() {
		_ = slowpoke.CloseAll()
	}()

	// file, _ := os.Open("knaben.xml")
	// defer file.Close()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)
	// feed, err := fp.Parse(file)
	if err != nil {
		panic(err)
	}
	// fmt.Println(feed.Title)

	handler.Parse(feed.Items)
}
