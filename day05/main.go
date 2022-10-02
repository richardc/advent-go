package main

import (
	_ "embed"
	"fmt"

	"advent2015/pkg/input"
	"advent2015/pkg/slices"
)

func isNice(s string) bool {
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
