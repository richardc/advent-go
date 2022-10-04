package day08

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   8,
		Input: func() any { return input.Lines(puzzle) },
		Part1: func(i any) any { return countSpecials(i.([]string)) },
		Part2: func(i any) any { return countEscaped(i.([]string)) },
	})
}

func countChars(s string) int {
	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '\\' {
			if s[i+1] == 'x' {
				// Will be \xFF
				i += 3
			} else {
				// \\ and \"
				i++
			}
		}
		count++
	}
	return count - 1
}

func countEscape(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' || s[i] == '"' {
			count++
		}
		count++
	}

	return count + 2
}

func countSpecial(s string) int {
	return len(s) - countChars(s)
}

func escapeGrowth(s string) int {
	return countEscape(s) - len(s)
}

func countSpecials(s []string) int {
	return slices.Sum(slices.Map(s, countSpecial))
}

func countEscaped(s []string) int {
	return slices.Sum(slices.Map(s, escapeGrowth))
}
