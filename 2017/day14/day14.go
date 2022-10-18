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
		hash := knothash.Hash([]byte(fmt.Sprintf("%s-%d", puzzle, row)))
		bits := strings.Join(slices.Map(hash[:], func(b byte) string { return fmt.Sprintf("%08b", b) }), "")
		total += len(slices.Filter([]byte(bits), func(c byte) bool { return c == byte('1') }))
	}
	return total
}
