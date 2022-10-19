package day16

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   16,
			Part1: func(any) any { return solve(strings.TrimSpace(puzzle)) },
			Part2: func(any) any { return solve2(strings.TrimSpace(puzzle)) },
		},
	)
}

func dance(floor string, moves []string) string {
	for _, move := range moves {
		var ia, ib int
		switch move[0] {
		case 's':
			pivot := len(floor) - input.MustAtoi(move[1:])
			floor = floor[pivot:] + floor[:pivot]
			continue
		case 'x':
			a, b, _ := strings.Cut(move[1:], "/")
			ia = input.MustAtoi(a)
			ib = input.MustAtoi(b)
		case 'p':
			a, b, _ := strings.Cut(move[1:], "/")
			ia = strings.Index(floor, a)
			ib = strings.Index(floor, b)
		}
		bfloor := []byte(floor)
		bfloor[ia], bfloor[ib] = bfloor[ib], bfloor[ia]
		floor = string(bfloor)
	}
	return floor
}

func solve(puzzle string) string {
	start := "abcdefghijklmnop"
	moves := strings.Split(puzzle, ",")
	return dance(start, moves)
}

func solve2(puzzle string) string {
	moves := strings.Split(puzzle, ",")
	start := "abcdefghijklmnop"
	floor := start

	iterations := 1_000_000_000
	for i := 0; i < iterations; i++ {
		floor = dance(floor, moves)
		if floor == start {
			period := i + 1
			skip := iterations / period
			i += period * (skip - 1)
		}
	}
	return floor
}
