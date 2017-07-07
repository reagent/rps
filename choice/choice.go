package choice

type Choice int

const (
	Rock Choice = iota
	Paper
	Scissors
)

func (c Choice) Beats(o Choice) bool {
	return (c == Rock && o == Scissors) || c > o
}
