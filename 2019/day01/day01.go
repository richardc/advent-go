package day01

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   1,
			Part1: func(any) any { return solve(puzzle) },
		},
	)
}

func moduleFuel(mass int) int {
	return (mass / 3) - 2
}

func solve(puzzle string) int {
	return slices.Sum(slices.Map(slices.Map(input.Lines(puzzle), input.MustAtoi), moduleFuel))
}
