package day01

import (
	"fmt"
	"strings"

	"github.com/mikkj17/aoc2018/utils"
	"github.com/mikkj17/aoc2018/utils/sets"
)

func parse(inp string) []int {
	return utils.Map(strings.Split(inp, "\n"), utils.ToInt)
}

func partOne(inp string) int {
	nums := parse(inp)
	return utils.Sum(nums)
}

func partTwo(inp string) int {
	nums := parse(inp)
	frequency := 0
	frequencies := sets.Empty[int]()
	for {
		for _, v := range nums {
			frequency += v
			if sets.Contains(frequencies, frequency) {
				return frequency
			}
			sets.Add(frequencies, frequency)
		}
	}
}

func Solve() {
	input := utils.ReadInput(1)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
