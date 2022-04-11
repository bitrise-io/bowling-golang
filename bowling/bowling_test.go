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
