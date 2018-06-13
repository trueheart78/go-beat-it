package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var name = "GameName"
var url = "https://howlongtobeat.com/game.php?id=12345"
var times = []string{"1.5", "2.0", "3.25", "--"}

// TestGameNewGame verifies that the NewGame method works as-expected
func TestGameNewGame(t *testing.T) {
	assert := assert.New(t)
	// assert equality
	g := NewGame(name, url, times...)
	assert.Equal(g.Name, name)
	assert.Equal(g.URL, url)
	assert.Equal(g.Main, times[0])
	assert.Equal(g.Extra, times[1])
	assert.Equal(g.Complete, times[2])
	assert.Equal(g.Average, times[3])
}
