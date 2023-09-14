package day02

import (
	"fmt"
	"strings"

	"github.com/mikkj17/aoc2018/utils"
)

func parse(inp string) []string {
	return strings.Split(inp, "\n")
}

func partOne(inp string) int {
	lines := parse(inp)
	counts := map[int]int{2: 0, 3: 0}

	for _, line := range lines {
		frequencies := utils.Counter([]rune(line))

		for _, count := range [2]int{2, 3} {
			if utils.ContainsValue(frequencies, count) {
				counts[count]++
			}
		}
	}

	return counts[2] * counts[3]
}

func getId(first string, second string) string {
	id := ""
	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			id += first[i : i+1]
		}
	}

	return id
}

func partTwo(inp string) string {
	lines := parse(inp)
	for i, first := range lines {
		for _, second := range lines[i+1:] {
			id := getId(first, second)
			if len(id) == len(first)-1 {
				return id
			}
		}
	}

	panic("wasn't found")
}

func Solve() {
	input := utils.ReadInput(2)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
