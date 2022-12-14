package day16

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   16,
		Input: func() any { return input.Lines(puzzle) },
		Part1: func(i any) any { return sueFinderPart1(i.([]string)) },
		Part2: func(i any) any { return sueFinderPart2(i.([]string)) },
	})
}

func matchSimple(s string, filter map[string]int) bool {
	// Sue 21: pomeranians: 7, trees: 0, goldfish: 10
	_, attrs, _ := strings.Cut(s, ": ")
	for _, thing := range strings.Split(attrs, ", ") {
		attr, value, _ := strings.Cut(thing, ": ")
		v := input.MustAtoi(value)
		if filter[attr] != v {
			return false
		}
	}
	return true
}

func sueFinder(sues []string, matcher func(string, map[string]int) bool) int {
	filter := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	for i, sue := range sues {
		if matcher(sue, filter) {
			return i + 1
		}
	}

	return 0
}

func sueFinderPart1(sues []string) int {
	return sueFinder(sues, matchSimple)
}

func matchTwo(s string, filter map[string]int) bool {
	// Sue 21: pomeranians: 7, trees: 0, goldfish: 10
	_, attrs, _ := strings.Cut(s, ": ")
	for _, thing := range strings.Split(attrs, ", ") {
		attr, value, _ := strings.Cut(thing, ": ")
		v := input.MustAtoi(value)

		switch attr {
		case "cats", "trees":
			if v <= filter[attr] {
				return false
			}
		case "pomeranians", "goldfish":
			if v >= filter[attr] {
				return false
			}
		default:
			if v != filter[attr] {
				return false
			}
		}
	}
	return true
}

func sueFinderPart2(sues []string) int {
	return sueFinder(sues, matchTwo)
}
