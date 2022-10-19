package day15

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
			Day:   15,
			Part1: func(any) any { return solve(puzzle) },
			Part2: func(any) any { return solve2(puzzle) },
		},
	)
}

func generatorA(state int) func() int {
	return func() int {
		state *= 16807
		state %= 2147483647
		return state
	}
}

func generatorB(state int) func() int {
	return func() int {
		state *= 48271
		state %= 2147483647
		return state
	}
}

func generatorA2(state int) func() int {
	return func() int {
		for {
			state *= 16807
			state %= 2147483647
			if state%4 == 0 {
				return state
			}
		}
	}
}

func generatorB2(state int) func() int {
	return func() int {
		for {
			state *= 48271
			state %= 2147483647
			if state%8 == 0 {
				return state
			}
		}
	}
}

func judge(genA, genB func() int, rounds int) int {
	matches := 0
	for i := 0; i < rounds; i++ {
		a := genA()
		b := genB()
		if a&0xFFFF == b&0xFFFF {
			matches++
		}
	}
	return matches
}

func solve(puzzle string) int {
	seeds := slices.Map(input.Lines(puzzle), func(s string) int {
		// Generator A starts with 883
		return input.MustAtoi(strings.Fields(s)[4])
	})
	return judge(generatorA(seeds[0]), generatorB(seeds[1]), 40_000_000)
}

func solve2(puzzle string) int {
	seeds := slices.Map(input.Lines(puzzle), func(s string) int {
		// Generator A starts with 883
		return input.MustAtoi(strings.Fields(s)[4])
	})
	return judge(generatorA2(seeds[0]), generatorB2(seeds[1]), 5_000_000)
}
