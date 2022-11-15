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
			Part2: func(any) any { return solve2(puzzle) },
		},
	)
}

func moduleFuel(mass int) int {
	return (mass / 3) - 2
}

func moduleFuelWithFuel(mass int) int {
	total := 0
	next := moduleFuel(mass)
	for next >= 0 {
		total += next
		next = moduleFuel(next)
	}
	return total
}

func solve(puzzle string) int {
	return slices.Sum(slices.Map(slices.Map(input.Lines(puzzle), input.MustAtoi), moduleFuel))
}

func solve2(puzzle string) int {
	return slices.Sum(slices.Map(slices.Map(input.Lines(puzzle), input.MustAtoi), moduleFuelWithFuel))
}
