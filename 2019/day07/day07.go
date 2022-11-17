package day07

import (
	_ "embed"

	"github.com/richardc/advent-go/2019/intcode"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   7,
			Part1: func(any) any { return maxSignal(puzzle) },
		},
	)
}

func maxSignal(puzzle string) int {
	program := intcode.NewCpu(puzzle)
	settings := []int{0, 1, 2, 3, 4}
	max := 0
	for _, sequence := range slices.Permutations(settings) {
		input := 0
		for i := 0; i < 5; i++ {
			stage := program.Clone()
			stage.Input([]int{sequence[i], input})
			stage.Run()
			input = stage.Output()[0]
		}
		if input > max {
			max = input
		}
	}
	return max
}
