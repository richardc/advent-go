package day24

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
			Year:  2017,
			Day:   24,
			Input: func() any { return makeChains(puzzle) },
			Part1: func(i any) any { return strongestBridge(i.([]Chain)) },
			Part2: func(i any) any { return longestStrongestBridge(i.([]Chain)) },
		},
	)
}

type Section struct {
	A, B int
}

func NewSection(s string) Section {
	astr, bstr, _ := strings.Cut(s, "/")
	a := input.MustAtoi(astr)
	b := input.MustAtoi(bstr)
	return Section{a, b}
}

func (s Section) Value() int {
	return s.A + s.B
}

func (s Section) String() string {
	return fmt.Sprintf("%d/%d", s.A, s.B)
}

type Chain struct {
	Links []Section
}

func (c Chain) String() string {
	return strings.Join(slices.Map(c.Links, Section.String), "--") + "\n"
}

func (c Chain) EndsWith() int {
	endsWith := 0
	for _, s := range c.Links {
		if s.A == endsWith {
			endsWith = s.B
		} else {
			endsWith = s.A
		}
	}
	return endsWith
}

func (c Chain) Append(s Section) Chain {
	return Chain{
		Links: append([]Section(nil), append(c.Links, s)...),
	}
}

func (c Chain) Next(sections []Section) []Chain {
	used := map[Section]struct{}{}
	for _, s := range c.Links {
		used[s] = struct{}{}
	}
	endsWidth := c.EndsWith()
	var links []Section
	for _, s := range sections {
		if _, ok := used[s]; ok {
			continue
		}
		if s.A == endsWidth || s.B == endsWidth {
			links = append(links, s)
		}
	}

	next := make([]Chain, len(links))
	for i, link := range links {
		next[i] = c.Append(link)
	}
	return next
}

func (c Chain) Strength() int {
	return slices.Sum(slices.Map(c.Links, Section.Value))
}

func (c Chain) Length() int {
	return len(c.Links)
}

func makeChains(puzzle string) []Chain {
	sections := slices.Map(input.Lines(puzzle), NewSection)
	var chains []Chain
	queue := []Chain{{}}
	for {
		var next []Chain
		for _, chain := range queue {
			next = append(next, chain.Next(sections)...)
		}
		if len(next) == 0 {
			break
		}
		chains = append(chains, next...)
		queue = next
	}
	return chains
}

func strongestBridge(chains []Chain) int {
	return slices.Max(slices.Map(chains, Chain.Strength))
}

func longestStrongestBridge(chains []Chain) int {
	maxLength := slices.Max(slices.Map(chains, Chain.Length))
	longChains := slices.Filter(chains, func(c Chain) bool { return maxLength == c.Length() })
	return strongestBridge(longChains)
}
