package day19

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"golang.org/x/exp/maps"
)

//go:embed input.txt
var puzzle string

type inp struct {
	m machine
	s string
}

func init() {
	runner.Register(runner.Solution{
		Day: 19,
		Input: func() any {
			lines := input.Lines(puzzle)
			return inp{
				m: newMachine(lines[:len(lines)-3]),
				s: lines[len(lines)-1],
			}
		},
		Part1: func(i any) any {
			in := i.(inp)
			return countOneReplacement(in.m, in.s)
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
