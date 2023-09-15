package day03

import (
	"github.com/mikkj17/aoc2018/utils"
	"github.com/mikkj17/aoc2018/utils/sets"
)

type Claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

func (c Claim) cover() sets.Set[utils.Position] {
	cover := sets.Empty[utils.Position]()
	for y := c.top; y < c.top+c.height; y++ {
		for x := c.left; x < c.left+c.width; x++ {
			sets.Add(cover, utils.Position{X: x, Y: y})
		}
	}

	return cover
}
