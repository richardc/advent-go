package day06

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
	"golang.org/x/exp/maps"
)

//go:embed "input.txt"
var puzzle string

func init() {
	runner.Register(
		runner.Solution{
			Year:  2019,
			Day:   6,
			Input: func() any { return newStarmap(puzzle) },
			Part1: func(i any) any { return i.(Starmap).Orbits() },
			Part2: func(i any) any { return i.(Starmap).Transfers() },
		},
	)
}

type Starmap struct {
	orbits map[string]string
}

func newStarmap(s string) Starmap {
	orbits := map[string]string{}

	for _, l := range input.Lines(s) {
		parent, child, _ := strings.Cut(l, ")")
		orbits[child] = parent
	}

	return Starmap{
		orbits: orbits,
	}
}

func (s Starmap) Orbit(body string) int {
	count := 0
	for {
		if parent, ok := s.orbits[body]; ok {
			count++
			body = parent
		} else {
			return count
		}
	}
}

func (s Starmap) Orbits() int {
	total := 0
	for body := range s.orbits {
		total += s.Orbit(body)
	}
	return total
}

func (s Starmap) Parents(body string) map[string]int {
	path := map[string]int{}
	steps := 0
	for {
		if parent, ok := s.orbits[body]; ok {
			body = parent
			path[body] = steps
			steps++
		} else {
			return path
		}
	}
}

func (s Starmap) Transfers() int {
	you := s.Parents("YOU")
	san := s.Parents("SAN")
	common := map[string]int{}
	for body, steps := range you {
		if santa, ok := san[body]; ok {
			common[body] = steps + santa
		}
	}
	return slices.Min(maps.Values(common))
}
