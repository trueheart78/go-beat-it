package parse

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/trueheart78/go-beat-it/internal/pkg/game"
)

func ExtractGames(html io.Reader) (games []game.Game, err error) {
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		log.Fatal(err)
	}

	// Find the games and their times
	doc.Find(".search_list_details").Each(func(i int, s *goquery.Selection) {
		name := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		d := s.Find(".center").Text()
		// fmt.Printf("[%v]\n", d)
		games = append(games, game.NewGame(name, url, extractTimes(d)...))
	})

	return
}

func extractTimes(hours string) []string {
	var times []string
	hours = strings.Replace(hours, " Hours", "", -1)
	hours = strings.Replace(hours, "½", ".5", -1)
	hours = strings.TrimSpace(hours)
	times = strings.Split(hours, " ")
	return times
}
