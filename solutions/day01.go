package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(inp string) []int {
	nums := []int{}

	for _, v := range strings.Split(inp, "\n") {
		value, _ := strconv.Atoi(v)
		nums = append(nums, value)
	}

	return nums
}

func partOne(inp string) int {
	nums := parse(inp)
	frequency := 0
	for _, v := range nums {
		frequency += v
	}

	return frequency
}

func partTwo(inp string) int {
	nums := parse(inp)
	frequency := 0
	frequencies := map[int]bool{}
	for {
		for _, v := range nums {
			frequency += v

			if frequencies[frequency] {
				return frequency
			}
			frequencies[frequency] = true
		}
	}
}

func dayOne() {
	data, _ := os.ReadFile("inputs/day01.txt")
	input := string(data)
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}
