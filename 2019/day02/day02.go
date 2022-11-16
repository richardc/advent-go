package day02

import (
	_ "embed"
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
			Year:  2019,
			Day:   2,
			Part1: func(any) any { return solve(puzzle) },
			Part2: func(any) any { return solve2(puzzle) },
		},
	)
}

func runIntcode(memory []int) int {
	i := 0
	for {
		switch memory[i] {
		case 1:
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
		case 2:
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
		case 99:
			return memory[0]
		}
		i += 4
	}
}

func solve(puzzle string) int {
	program := slices.Map(strings.Split(strings.Trim(puzzle, "\n"), ","), input.MustAtoi)
	program[1] = 12
	program[2] = 2
	return runIntcode(program)
}

func solve2(puzzle string) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			program := slices.Map(strings.Split(strings.Trim(puzzle, "\n"), ","), input.MustAtoi)
			program[1] = i
			program[2] = j
			if runIntcode(program) == 19690720 {
				return i*100 + j
			}
		}
	}
	return -1
}
