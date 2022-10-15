package day02

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   2,
			Part1: func(any) any { return solve(puzzle) },
		},
	)
}

func checksum(line string) int {
	min, max := slices.MinMax(slices.Map(strings.Fields(line), input.MustAtoi))
	return max - min
}

func solve(puzzle string) int {
	return slices.Sum(slices.Map(input.Lines(puzzle), checksum))
}
