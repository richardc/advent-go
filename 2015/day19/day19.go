package day19

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day: 19,
		Part1: func(any) any {
			lines := input.Lines(puzzle)
			m := newMachine(lines[:len(lines)-3])
			s := lines[len(lines)-1]
			return countOneReplacement(m, s)
		},
		Part2: func(any) any {
			lines := input.Lines(puzzle)
			m := newReverseMachine(lines[:len(lines)-3])
			s := lines[len(lines)-1]
			return countStepsTillExpanded(m, s, "e")
		},
	})
}

type machine map[string][]string

func newMachine(s []string) machine {
	// H => HO
	m := machine{}
	for _, pair := range s {
		left, right, _ := strings.Cut(pair, " => ")
		m[left] = append(m[left], right)
	}
	return m
}

func newReverseMachine(s []string) machine {
	m := machine{}
	for _, pair := range s {
		left, right, _ := strings.Cut(pair, " => ")
		m[right] = append(m[right], left)
	}
	return m
}

func countOneReplacement(m machine, s string) int {
	expansions := map[string]struct{}{}
	for k, v := range m {
		start := 0
		for {
			i := strings.Index(s[start:], k)
			if i == -1 {
				break
			}
			for _, replace := range v {
				exp := s[:start+i] + replace + s[start+len(k)+i:]
				expansions[exp] = struct{}{}
			}
			start = start + i + 1
		}
	}
	return len(maps.Keys(expansions))
}

func countStepsTillExpanded(m machine, start, end string) int {
	// Try the machine rules longest-first
	rules := maps.Keys(m)
	slices.SortFunc(rules, func(a, b string) bool { return len(b) < len(a) })

	steps := 0
	input := start
	for {
		for _, rule := range rules {
			i := strings.Index(input, rule)
			if i == -1 {
				continue
			}
			input = input[:i] + m[rule][0] + input[len(rule)+i:]
			steps++
			if input == end {
				return steps
			}
		}
	}
}
