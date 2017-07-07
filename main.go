package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/reagent/rps/game"
	"github.com/reagent/rps/players"
)

const DefaultRounds int = 100

func main() {
	ps := players.Players{
		// &players.PaperPlayer{},
		// &players.RockPlayer{},
		// &players.ScissorsPlayer{},
	}

	if len(ps) < 2 {
		fmt.Println("Must supply 2 or more players")
		os.Exit(1)
	}

	rounds := flag.Int("rounds", DefaultRounds, "number of rounds to play")

	flag.Parse()

	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			g := game.New(ps[i], ps[j])
			g.Play(*rounds)
			g.Print()
		}
	}
}
