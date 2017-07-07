package game

import (
	"fmt"
	"reflect"

	"github.com/reagent/rps/players"
	"github.com/reagent/rps/result"
)

type Game struct {
	p1 players.Player
	p2 players.Player
	s1 float64
	s2 float64
}

func New(p1 players.Player, p2 players.Player) *Game {
	return &Game{p1: p1, p2: p2, s1: 0.0, s2: 0.0}
}

func (g *Game) Play(rounds int) {
	g.p1.Initialize(g.p2)
	g.p2.Initialize(g.p1)

	for i := 0; i < rounds; i++ {
		c1, c2 := g.p1.Choose(), g.p2.Choose()

		if c1.Beats(c2) {
			g.s1 += 1.0

			g.p1.Result(c1, c2, result.Win)
			g.p2.Result(c2, c1, result.Lose)
		} else if c2.Beats(c1) {
			g.s2 += 1.0

			g.p1.Result(c1, c2, result.Lose)
			g.p2.Result(c2, c1, result.Win)
		} else {
			g.s1 += 0.5
			g.s2 += 0.5

			g.p1.Result(c1, c2, result.Draw)
			g.p2.Result(c2, c1, result.Draw)
		}
	}
}

func (g *Game) Print() {
	fmt.Printf("%s vs. %s\n", name(g.p1), name(g.p2))
	fmt.Printf("\t%s: %.1f\n", name(g.p1), g.s1)
	fmt.Printf("\t%s: %.1f\n", name(g.p2), g.s2)

	if g.s1 == g.s2 {
		fmt.Printf("\tDraw\n")
	} else if g.s1 > g.s2 {
		fmt.Printf("\t%s Wins\n", name(g.p1))
	} else {
		fmt.Printf("\t%s Wins\n", name(g.p2))
	}

	fmt.Println()
}

func name(p players.Player) string {
	return reflect.TypeOf(p).Elem().Name()
}
