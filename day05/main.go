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
	for i, _ := range s[:len(s)-1] {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func isReallyNice(s string) bool {
	return framedChar(s) && doubleDouble(s)
}

func doubleDouble(s string) bool {
	for i, _ := range s[:len(s)-2] {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func framedChar(s string) bool {
	for i, _ := range s[:len(s)-2] {
		if s[i] == s[i+2] {
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
	part2 := len(slices.Filter(input.Lines(puzzle), isReallyNice))
	fmt.Println("Part 2", part2)
}
