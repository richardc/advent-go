package day17

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"golang.org/x/exp/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   17,
			Part1: func(any) any { return valueAfter2017(puzzle) },
		},
	)
}

func spinlock(step, iterations int) []int {
	memory := []int{0}
	position := 0
	for i := 0; i < iterations; i++ {
		position = (position+step)%len(memory) + 1
		memory = slices.Insert(memory, position, i+1)
	}
	return memory
}

func valueAfter2017(puzzle string) int {
	step := input.MustAtoi(strings.TrimSpace(puzzle))
	memory := spinlock(step, 2018)
	index := slices.Index(memory, 2017) + 1
	return memory[index]
}
