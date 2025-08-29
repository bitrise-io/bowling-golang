package bowling_test

import (
	"testing"

	"github.com/bitrise-io/bowling-golang/bowling"
)

func TestScore(t *testing.T) {
	testCases := []struct {
		desc    string
		rollAll func(g *bowling.Game)
		want    int
	}{
		{
			desc: "gutter",
			rollAll: func(g *bowling.Game) {
				g.Roll(0)
				g.Roll(0)

				g.Roll(0)
				g.Roll(0)

				g.Roll(0)
				g.Roll(0)
			},
			want: 0,
		},
		{
			desc: "simple game",
			rollAll: func(g *bowling.Game) {
				g.Roll(3)
				g.Roll(2)

				g.Roll(4)
				g.Roll(1)

				g.Roll(2)
				g.Roll(3)
			},
			want: 15,
		},
		{
			desc: "spare",
			rollAll: func(g *bowling.Game) {
				g.Roll(7)
				g.Roll(3) // spare

				g.Roll(4) // bonus for spare
				g.Roll(2)

				g.Roll(1)
				g.Roll(0)
			},
			want: 21, // (7+3+4) + (4+2) + (1+0) = 14 + 6 + 1 = 21
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			g := bowling.Game{}
			tc.rollAll(&g)
			got := g.Score()
			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
