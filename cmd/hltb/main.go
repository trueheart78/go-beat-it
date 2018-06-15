package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/trueheart78/go-beat-it/internal/pkg/game"
	"github.com/trueheart78/go-beat-it/internal/pkg/parse"
	"github.com/trueheart78/go-beat-it/internal/pkg/search"
)

const version = 1.0

func main() {
	var err error
	var html io.Reader
	var games []game.Game
	var gameName = flag.String("g", "", "game name to lookup")
	var versionCheck = flag.Bool("v", false, "version")

	flag.Parse()

	if *versionCheck {
		fmt.Printf("version %0.2f\n", version)

		os.Exit(0)
	}

	if *gameName == "" {
		flag.Usage()
		os.Exit(1)
	}

	html, err = search.Submit(*gameName, search.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	games, err = parse.ExtractGames(html)
	if err != nil {
		log.Fatal(err)
	}

	if len(games) > 0 {
		for i, g := range games {
			fmt.Printf("%d. %v\n", (i + 1), g.ToString())
		}
	} else {
		fmt.Println("No games found")
		os.Exit(1)
	}
}
