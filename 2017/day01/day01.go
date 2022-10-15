package day01

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   1,
			Part1: func(any) any { return uncaptcha(strings.TrimSpace(puzzle)) },
			Part2: func(any) any { return uncaptcha_half(strings.TrimSpace(puzzle)) },
		},
	)
}

func value(c byte) byte {
	return c - byte('0')
}

func uncaptcha(puzzle string) int {
	sum := 0
	for i := range puzzle {
		if puzzle[i] == puzzle[(i+1)%len(puzzle)] {
			sum += int(value(puzzle[i]))
		}
	}
	return sum
}

func uncaptcha_half(puzzle string) int {
	sum := 0
	for i := range puzzle {
		if puzzle[i] == puzzle[(i+len(puzzle)/2)%len(puzzle)] {
			sum += int(value(puzzle[i]))
		}
	}
	return sum
}
