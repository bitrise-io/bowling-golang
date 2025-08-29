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
func (g *Game) Score() int {
	score := 0
	for _, pins := range g.rolls {
		score += pins
	}
	return score
}
