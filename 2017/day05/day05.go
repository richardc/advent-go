package day05

import (
	_ "embed"

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
			Day:   5,
			Part1: func(any) any { return jumps(slices.Map(input.Lines(puzzle), input.MustAtoi)) },
			Part2: func(any) any { return jumpsOddly(slices.Map(input.Lines(puzzle), input.MustAtoi)) },
		},
	)
}

func jumps(program []int) int {
	pc := 0
	steps := 0
	for pc >= 0 && pc < len(program) {
		steps++
		next := pc + program[pc]
		program[pc]++
		pc = next
	}
	return steps
}

func jumpsOddly(program []int) int {
	pc := 0
	steps := 0
	for pc >= 0 && pc < len(program) {
		steps++
		next := pc + program[pc]
		if program[pc] >= 3 {
			program[pc]--
		} else {
			program[pc]++
		}
		pc = next
	}
	return steps
}
