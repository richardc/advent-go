package day05

import (
	_ "embed"

	"github.com/richardc/advent-go/2019/intcode"
	"github.com/richardc/advent-go/runner"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   5,
			Part1: func(any) any { return solve(puzzle, 1) },
			Part2: func(any) any { return solve(puzzle, 5) },
		},
	)
}

func solve(puzzle string, input int) int {
	cpu := intcode.New(puzzle)
	cpu.Input([]int{input})
	cpu.Run()
	output := cpu.Output()
	return output[len(output)-1]
}
