package day17

import (
	_ "embed"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   17,
		Input: func() any { return slices.Map(input.Lines(puzzle), input.MustAtoi) },
		Part1: func(i any) any { return containers(i.([]int), 150) },
		Part2: func(i any) any { return smallContainers(i.([]int), 150) },
	})
}

func containers(containers []int, eggnog int) int {
	count := 0
	for _, set := range slices.Powerset(containers) {
		if eggnog == slices.Sum(set) {
			count++
		}
	}
	return count
}

func smallContainers(containers []int, eggnog int) int {
	counts := map[int]int{}
	for _, set := range slices.Powerset(containers) {
		if eggnog == slices.Sum(set) {
			counts[len(set)]++
		}
	}

	return counts[slices.Min(maps.Keys(counts))]
}
