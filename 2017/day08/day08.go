package day08

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/maths"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year: 2017,
			Day:  8,
			Input: func() any {
				terminal, running := maxRegister(puzzle)
				return []int{terminal, running}
			},
			Part1: func(i any) any { return i.([]int)[0] },
			Part2: func(i any) any { return i.([]int)[1] },
		},
	)
}

func maxRegister(puzzle string) (terminal int, running int) {
	registers := map[string]int{}
	for _, instr := range input.Lines(puzzle) {
		toks := strings.Fields(instr)
		check := registers[toks[4]]
		value := input.MustAtoi(toks[6])

		matched := false
		switch toks[5] {
		case ">":
			matched = check > value
		case ">=":
			matched = check >= value
		case "<":
			matched = check < value
		case "<=":
			matched = check <= value
		case "==":
			matched = check == value
		case "!=":
			matched = check != value
		}

		if !matched {
			continue
		}

		switch toks[1] {
		case "inc":
			registers[toks[0]] += input.MustAtoi(toks[2])
		case "dec":
			registers[toks[0]] -= input.MustAtoi(toks[2])
		}

		running = maths.Max(running, registers[toks[0]])
	}
	return slices.Max(maps.Values(registers)), running
}
