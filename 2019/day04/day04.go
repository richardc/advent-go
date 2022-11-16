package day04

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   4,
			Part1: func(any) any { return solve(puzzle) },
			Part2: func(any) any { return solve2(puzzle) },
		},
	)
}

func doubleChar(password string) bool {
	for i := 0; i < 5; i++ {
		if password[i] == password[i+1] {
			return true
		}
	}
	return false
}

func justDoubleChar(password string) bool {
	groups := slices.GroupBy([]byte(password))
	for _, group := range groups {
		if len(group.Group) == 2 {
			return true
		}
	}
	return false
}

func increasing(password string) bool {
	for i := 0; i < 5; i++ {
		if password[i] > password[i+1] {
			return false
		}
	}
	return true
}

func validPassword(password string) bool {
	return doubleChar(password) && increasing(password)
}

func solve(puzzle string) int {
	mins, maxs, _ := strings.Cut(strings.Trim(puzzle, "\n"), "-")
	min := input.MustAtoi(mins)
	max := input.MustAtoi(maxs)
	count := 0
	for i := min; i < max; i++ {
		if validPassword(fmt.Sprintf("%d", i)) {
			count++
		}
	}
	return count
}

func validPassword2(password string) bool {
	return justDoubleChar(password) && increasing(password)
}

func solve2(puzzle string) int {
	mins, maxs, _ := strings.Cut(strings.Trim(puzzle, "\n"), "-")
	min := input.MustAtoi(mins)
	max := input.MustAtoi(maxs)
	count := 0
	for i := min; i < max; i++ {
		if validPassword2(fmt.Sprintf("%d", i)) {
			count++
		}
	}
	return count
}
