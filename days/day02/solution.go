package day02

import (
	"fmt"
	"os"
	"strings"
)

func parse(inp string) []string {
	return strings.Split(inp, "\n")
}

func containsValue(frequencies map[rune]int, value int) bool {
	for _, count := range frequencies {
		if count == value {
			return true
		}
	}

	return false
}

func makeFrequencies(line string) map[rune]int {
	frequencies := map[rune]int{}
	for _, char := range line {
		frequencies[char]++
	}

	return frequencies
}

func partOne(inp string) int {
	lines := parse(inp)
	counts := map[int]int{2: 0, 3: 0}

	for _, line := range lines {
		frequencies := makeFrequencies(line)

		if containsValue(frequencies, 2) {
			counts[2]++
		}
		if containsValue(frequencies, 3) {
			counts[3]++
		}
	}

	return counts[2] * counts[3]
}

func getId(first string, second string) string {
	id := ""
	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			id += string(first[i])
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
	data, _ := os.ReadFile("days/day02/input.txt")
	input := string(data)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
