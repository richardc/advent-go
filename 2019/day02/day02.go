package day02

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
			Day:   2,
			Part1: func(any) any { return solve(puzzle) },
			Part2: func(any) any { return solve2(puzzle) },
		},
	)
}

func solve(puzzle string) int {
	program := intcode.NewCpu(puzzle)
	program.Set(1, 12)
	program.Set(2, 2)
	program.Run()
	return program.Get(0)
}

func solve2(puzzle string) int {
	program := intcode.NewCpu(puzzle)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			run := program.Clone()
			run.Set(1, i)
			run.Set(2, j)
			run.Run()
			if run.Get(0) == 19690720 {
				return i*100 + j
			}
		}
	}
	return -1
}
