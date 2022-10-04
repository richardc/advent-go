package dayXX

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
		Day:   99,
		Input: func() any { return input.Lines(puzzle) },
		Part1: func(i any) any { return sumXXX(i.([]string)) },
	})
}

func XXX(s string) int {
	return 0
}

func sumXXX(s []string) int {
	return slices.Sum(slices.Map(s, XXX))
}
