package day24

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   24,
		Input: func() any { return slices.Map(input.Lines(puzzle), input.MustAtoi) },
		Part1: func(i any) any { return smallestQE(i.([]int)) },
	})
}

func smallestQE(s []int) int {
	goal := slices.Sum(s) / 3
	for n := 1; n < len(s); n++ {
		// fmt.Printf("Testing with a basket of %d from %d\n", n, len(s))
		var fits [][]int
		slices.CombinationsFunc(s, n, func(s []int) {
			if goal == slices.Sum(s) {
				fits = append(fits, append([]int{}, s...))
			}
		})
		if len(fits) == 0 {
			continue
		}
		return slices.Min(slices.Map(fits, slices.Product[int]))
	}
	return 0
}
