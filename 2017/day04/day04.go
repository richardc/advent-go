package day04

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	slcs "github.com/richardc/advent-go/slices"
	"golang.org/x/exp/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   4,
			Part1: func(any) any { return part1(puzzle) },
			Part2: func(any) any { return part2(puzzle) },
		},
	)
}

func nothing_repeats(s string) bool {
	seen := map[string]struct{}{}
	for _, word := range strings.Fields(s) {
		if _, ok := seen[word]; ok {
			return false
		}
		seen[word] = struct{}{}
	}
	return true
}

func part1(puzzle string) int {
	return len(slcs.Filter(input.Lines(puzzle), nothing_repeats))
}

func nothing_palindromes(s string) bool {
	seen := map[string]struct{}{}
	for _, word := range strings.Fields(s) {
		b := []byte(word)
		slices.Sort(b)
		word := string(b)
		if _, ok := seen[word]; ok {
			return false
		}
		seen[word] = struct{}{}
	}
	return true
}

func part2(puzzle string) int {
	return len(slcs.Filter(input.Lines(puzzle), nothing_palindromes))
}
