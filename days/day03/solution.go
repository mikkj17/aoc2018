package day03

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mikkj17/aoc2018/utils"
	"github.com/mikkj17/aoc2018/utils/sets"
)

func parse(inp string) []Claim {
	r := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)
	return utils.Map(strings.Split(inp, "\n"), func(line string) Claim {
		groups := r.FindStringSubmatch(line)
		vals := utils.Map(groups[1:], func(s string) int {
			return utils.ToInt(s)
		})
		return Claim{
			id:     vals[0],
			left:   vals[1],
			top:    vals[2],
			width:  vals[3],
			height: vals[4],
		}
	})
}

func partOne(inp string) int {
	claims := parse(inp)
	cover := []utils.Position{}
	for _, c := range claims {
		cover = append(cover, sets.ToSlice(c.cover())...)
	}
	counts := utils.Counter(cover)

	twoOrMore := 0
	for _, count := range counts {
		if count >= 2 {
			twoOrMore++
		}
	}

	return twoOrMore
}

func partTwo(inp string) int {
	claims := parse(inp)
	for i, claim := range claims {
		overlap := false
		for j, c := range claims {
			if i != j {
				intersect := sets.Intersection(claim.cover(), c.cover())
				if sets.Length(intersect) != 0 {
					overlap = true
					break
				}
			}
		}
		if !overlap {
			return claim.id
		}

	}

	panic("wasn't found")
}

func Solve() {
	input := utils.ReadInput(3)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
