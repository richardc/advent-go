package day01

import (
	_ "embed"

	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var input string

func init() {
	runner.Register(runner.Solution{
		Day:   1,
		Part1: func(i any) any { return whatFloor(input) },
		Part2: func(i any) any { return goesNegative(input) },
	})
}

func whatFloor(input string) int {
	count := 0
	for _, c := range input {
		switch c {
		case '(':
			count++
		case ')':
			count--
		}
	}
	return count
}

func goesNegative(input string) int {
	count := 0
	for i, c := range input {
		switch c {
		case '(':
			count++
		case ')':
			count--
		}

		if count < 0 {
			return i + 1
		}
	}
	return -1
}
