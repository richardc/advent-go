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
			Part2: func(any) any { return maxLoopedSignal(puzzle) },
		},
	)
}

func maxSignal(puzzle string) int {
	program := intcode.New(puzzle)
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

func maxLoopedSignal(puzzle string) int {
	program := intcode.New(puzzle)
	settings := []int{5, 6, 7, 8, 9}
	max := 0

	cpus := slices.Map(settings, func(i int) *intcode.Cpu {
		cpu := program.Clone()
		return &cpu
	})
	for i := 0; i < 4; i++ {
		next := i + 1
		cpus[i].OutputFunc(func(v int) { cpus[next].Input([]int{v}) })
	}
	cpus[4].OutputFunc(func(v int) {
		cpus[0].Input([]int{v})
		if v > max {
			max = v
		}
	})
	for _, sequence := range slices.Permutations(settings) {
		for i, cpu := range cpus {
			cpu.Reset(&program)
			cpu.Input([]int{sequence[i]})
		}
		cpus[0].Input([]int{0})

		for {
			for _, cpu := range cpus {
				cpu.Run()
			}

			if slices.All(cpus, func(c *intcode.Cpu) bool { return c.Halted() }) {
				break
			}
		}
	}
	return max
}
