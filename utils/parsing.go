package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ToInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return val
}

func ReadInput(day int) string {
	path := fmt.Sprintf("days/day%02d/input.txt", day)
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}
