package day06

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   6,
			Part1: func(any) any { return steps(puzzle) },
			Part2: func(any) any { return cycle(puzzle) },
		},
	)
}

func rebalance(memory []int) (int, int) {
	seen := map[string]int{}
	step := 1

	for {
		// Find first biggest pile
		max, index := 0, 0
		for i, v := range memory {
			if v > max {
				max = v
				index = i
			}
		}

		distribute := memory[index]
		memory[index] = 0

		for distribute > 0 {
			distribute--
			index++
			index %= len(memory)
			memory[index]++
		}

		key := strings.Join(slices.Map(memory, strconv.Itoa), ", ")
		if when, ok := seen[key]; ok {
			return step, step - when
		}
		seen[key] = step
		step++
	}
}

func steps(puzzle string) int {
	steps, _ := rebalance(slices.Map(strings.Fields(puzzle), input.MustAtoi))
	return steps
}

func cycle(puzzle string) int {
	_, cycle := rebalance(slices.Map(strings.Fields(puzzle), input.MustAtoi))
	return cycle
}
