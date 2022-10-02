package main

import (
	_ "embed"
	"fmt"
	"strings"

	"advent2015/pkg/input"
	"advent2015/pkg/slices"
)

func isNice(s string) bool {
	switch {
	case
		strings.Contains(s, "ab"),
		strings.Contains(s, "cd"),
		strings.Contains(s, "pq"),
		strings.Contains(s, "xy"):
		return false
	}

	vowels := strings.Count(s, "a") +
		strings.Count(s, "e") +
		strings.Count(s, "i") +
		strings.Count(s, "o") +
		strings.Count(s, "u")

	if vowels < 3 {
		return false
	}

	return doubledChar(s)
}

func doubledChar(s string) bool {
	bytes := []byte(s)
	for i, c := range bytes[:len(bytes)-1] {
		if c == s[i+1] {
			return true
		}
	}
	return false
}

//go:embed input.txt
var puzzle string

func main() {
	part1 := len(slices.Filter(input.Lines(puzzle), isNice))
	fmt.Println("Part 1", part1)
	part2 := 0
	fmt.Println("Part 2", part2)
}
