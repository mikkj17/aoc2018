package day05

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mikkj17/aoc2018/utils"
	"github.com/mikkj17/aoc2018/utils/sets"
)

func isSame(a rune, b rune) bool {
	if unicode.IsUpper(a) {
		return unicode.IsLower(b) && a == unicode.ToUpper(b)
	} else {
		return unicode.IsUpper(b) && a == unicode.ToLower(b)
	}
}

func react(polymer string) string {
	for {
		breakOut := true
		for i, a := range polymer[:len(polymer)-1] {
			b := polymer[i+1]
			if isSame(a, rune(b)) {
				polymer = polymer[:i] + polymer[i+2:]
				breakOut = false
				break
			}
		}

		if breakOut {
			break
		}
	}

	return polymer
}

func partOne(inp string) int {
	return len(react(inp))
}

func partTwo(inp string) int {
	distinctChars := sets.FromSlice([]rune(strings.ToLower(inp)))
	diffs := map[rune]int{}
	for c := range distinctChars.Values {
		polymer := strings.ReplaceAll(inp, string(c), "")
		polymer = strings.ReplaceAll(polymer, string(unicode.ToUpper(c)), "")
		diffs[c] = len(react(polymer))
	}

	_, length := utils.MinValue(diffs)
	return length
}

func Solve() {
	input := utils.ReadInput(5)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
