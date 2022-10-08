package day20

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"golang.org/x/exp/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   20,
		Part1: func(any) any { return lowestHouse(input.MustAtoi(strings.TrimSpace(puzzle))) },
	})
}

func lowestHouse(target int) int {
	max := target / 10
	presents := make([]int, max)
	for elf := 1; elf < max; elf++ {
		for house := elf; house < max; house += elf {
			presents[house] += elf * 10
		}
	}

	return slices.IndexFunc(presents, func(i int) bool { return i >= target })
}
