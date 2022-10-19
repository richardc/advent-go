package day17

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"golang.org/x/exp/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2017,
			Day:   17,
			Part1: func(any) any { return valueAfter2017(puzzle) },
			Part2: func(any) any { return valueAfter50Mil(puzzle) },
		},
	)
}

func spinlock(step, iterations int) []int {
	memory := []int{0}
	position := 0
	for i := 0; i < iterations; i++ {
		position = (position+step)%len(memory) + 1
		memory = slices.Insert(memory, position, i+1)
	}
	return memory
}

func valueAfter2017(puzzle string) int {
	step := input.MustAtoi(strings.TrimSpace(puzzle))
	memory := spinlock(step, 2018)
	index := slices.Index(memory, 2017) + 1
	return memory[index]
}

func valueAfter0After(step, count int) int {
	memory := spinlock(step, count)
	index := slices.Index(memory, 0) + 1
	return memory[index]
}

func InspectValueAfter0After(puzzle string) int {
	step := input.MustAtoi(strings.TrimSpace(puzzle))
	for i := 1; i < 1000; i++ {
		fmt.Printf("%d %d\n", i, valueAfter0After(step, i))
	}
	return -1
}

func valueAfter50Mil(puzzle string) int {
	step := input.MustAtoi(strings.TrimSpace(puzzle))
	value := 0
	position := 0
	for i := 1; i < 50_000_001; i++ {
		position = (position + step + 1) % i
		if position == 0 {
			// insert at the start, new neighbour for 0
			value = i
		}
	}
	return value
}
