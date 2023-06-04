package day13

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
			Day:   13,
			Part1: func(any) any { return solve(puzzle) },
		},
	)
}

func solve(puzzle string) int {
	cpu := intcode.New(puzzle)
	cpu.Run()

	blocks := 0
	for i, val := range cpu.Output() {
		if i%3 == 2 {
			if val == 2 {
				blocks++
			}
		}
	}

	return blocks
}
