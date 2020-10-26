package main

import (
	"fmt"

	"github.com/bitrise-io/bowling-golang/bowling"
)

func main() {
	{
		// simple game
		g := bowling.Game{}

		g.Roll(2)
		g.Roll(3)

		g.Roll(2)
		g.Roll(3)

		fmt.Println(g.Score()) // 10
	}

	{
		// spare
		g := bowling.Game{}

		g.Roll(8)
		g.Roll(2)

		g.Roll(5)
		g.Roll(1)

		fmt.Println(g.Score()) // 21
	}

	{
		// strike
		g := bowling.Game{}

		g.Roll(10)

		g.Roll(5)
		g.Roll(1)

		fmt.Println(g.Score()) // 22
	}
}
