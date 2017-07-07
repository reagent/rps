package players

import (
	"github.com/reagent/rps/choice"
	"github.com/reagent/rps/result"
)

type Player interface {
	Initialize(Player)
	Choose() choice.Choice
	Result(choice.Choice, choice.Choice, result.Result)
}

type Players []Player

type BasicPlayer struct{}

func (p *BasicPlayer) Initialize(Player)                                  {}
func (p *BasicPlayer) Result(choice.Choice, choice.Choice, result.Result) {}
