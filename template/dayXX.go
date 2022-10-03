package dayXX

import (
	_ "embed"

	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var input string

func init() {
	runner.Register(runner.Solution{
		Day:   0,
		Part1: func(any) any { return XXX(input) },
	})
}

func XXX(s string) int {
	return 0
}
