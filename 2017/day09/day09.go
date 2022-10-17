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
			Part2: func(any) any { return garbage(puzzle) },
		},
	)
}

func score(puzzle string) int {
	score, _ := scan(puzzle)
	return score
}

func garbage(puzzle string) int {
	_, garbage := scan(puzzle)
	return garbage
}

func scan(puzzle string) (score int, garbage int) {
	current := 0
	junk := false
	for i := 0; i < len(puzzle); i++ {
		switch puzzle[i] {
		case '{':
			if !junk {
				current++
				continue
			}
		case '}':
			if !junk {
				score += current
				current--
				continue
			}
		case '!':
			i++
			continue
		case '<':
			if !junk {
				junk = true
				continue
			}
		case '>':
			junk = false
			continue
		}

		if junk {
			garbage++
		}

	}
	return score, garbage
}
