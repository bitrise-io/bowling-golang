package bowling

// Game ...
type Game struct {
	rolls []int
}

// Roll ...
func (g *Game) Roll(n int) {
	g.rolls = append(g.rolls, n)
}

// Score ...
func (g *Game) Score() (score int) {
	return 0
}
