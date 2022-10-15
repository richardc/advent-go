package day02

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
			Day:   2,
			Input: func() any { return input.Sheet(puzzle, input.MustAtoi) },
			Part1: func(i any) any { return checksums(i.([][]int)) },
		},
	)
}

func checksum(row []int) int {
	min, max := slices.MinMax(row)
	return max - min
}

func checksums(sheet [][]int) int {
	return slices.Sum(slices.Map(sheet, checksum))
}
