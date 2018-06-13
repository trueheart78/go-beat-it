package parse

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExtractGames(t *testing.T) {
	assert := assert.New(t)

	html, err := os.Open("test_data/dark-souls.html")
	if err != nil {
		log.Fatal(err)
	}

	games, _ := ExtractGames(html)
	assert.Equal(15, len(games))

	html, err = os.Open("test_data/no-results.html")
	if err != nil {
		log.Fatal(err)
	}
	games, _ = ExtractGames(html)
	assert.Equal(0, len(games))

	html, err = os.Open("test_data/limited-details.html")
	if err != nil {
		log.Fatal(err)
	}
	games, _ = ExtractGames(html)
	assert.Equal(1, len(games))
	assert.Equal("Final Fantasy X/X-2 HD Remaster", games[0].Name)
}

func TestParseExtractTimes(t *testing.T) {
	received := extractTimes("48½ Hours 62½ Hours 108 Hours 62 Hours ")
	expected := []string{"48.5", "62.5", "108", "62"}

	assert.Equal(t, expected, received)

	received = extractTimes("46½ Hours -- --")
	expected = []string{"46.5", "--", "--"}

	assert.Equal(t, expected, received)
}
