package day09

import (
	_ "embed"

	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   9,
			Part1: func(any) any { return score(puzzle) },
		},
	)
}

func score(puzzle string) int {
	total := 0
	current := 0
	junk := false
	for i := 0; i < len(puzzle); i++ {
		switch puzzle[i] {
		case '{':
			if !junk {
				current++
			}
		case '}':
			if !junk {
				total += current
				current--
			}
		case '!':
			i++
		case '<':
			junk = true
		case '>':
			junk = false
		}
	}
	return total
}
