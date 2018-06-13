// Package game makes a Game struct available
package game

// Game contains details of how long to beat
type Game struct {
	Name     string
	URL      string
	Main     string
	Extra    string
	Complete string
	Average  string
}

func (g Game) ToString() string {

	return ""
}

// NewGame constructs a new game
func NewGame(name string, url string, times ...string) Game {
	g := Game{Name: name, URL: url}
	if len(times) >= 1 {
		g.set("Main", times[0])
	}
	if len(times) >= 2 {
		g.set("Extra", times[1])
	}
	if len(times) >= 3 {
		g.set("Complete", times[2])
	}
	if len(times) >= 4 {
		g.set("Average", times[3])
	}
	return g
}

func (g *Game) set(name string, time string) {
	switch name {
	case "Main":
		g.Main = time
	case "Extra":
		g.Extra = time
	case "Complete":
		g.Complete = time
	case "Average":
		g.Average = time
	}
}
