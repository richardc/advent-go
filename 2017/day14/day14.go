package day14

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/richardc/advent-go/2017/knothash"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   14,
			Part1: func(any) any { return countSet(strings.TrimSpace(puzzle)) },
		},
	)
}

func countSet(puzzle string) int {
	total := 0
	for row := 0; row < 128; row++ {
		bits := fmt.Sprintf("%b", knothash.Hash([]byte(fmt.Sprintf("%s-%d", puzzle, row))))
		total += len(slices.Filter([]byte(bits), func(c byte) bool { return c == byte('1') }))
	}
	return total
}
