package day13

import (
	_ "embed"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/math"
	"github.com/richardc/advent-go/runner"
	"github.com/richardc/advent-go/slices"
)

//go:embed input.txt
var puzzle string

func init() {
	runner.Register(runner.Solution{
		Day:   13,
		Input: func() any { return input.Lines(puzzle) },
		Part1: func(i any) any { return happiestSeating(i.([]string)) },
	})
}

func edge(a, b string) [2]string {
	if a < b {
		return [2]string{a, b}
	} else {
		return [2]string{b, a}
	}
}

func happiestSeating(s []string) int {
	graph := map[[2]string]int{}
	names := map[string]struct{}{}
	for _, r := range s {
		// "Alice would gain 27 happiness units by sitting next to Bob."
		toks := strings.Fields(r)
		guest1 := toks[0]
		guest2 := toks[10]
		guest2 = guest2[:len(guest2)-1] // drop the last '.'

		names[guest1] = struct{}{}
		names[guest2] = struct{}{}

		gain, _ := strconv.Atoi(toks[3])
		if toks[2] == "lose" {
			gain *= -1
		}

		graph[edge(guest1, guest2)] += gain
	}

	guests := maps.Keys(names)
	arrangements := slices.Permutations(guests)

	max := 0
	for _, seating := range arrangements {
		value := graph[edge(seating[0], seating[len(seating)-1])]
		for i := 0; i < len(seating)-1; i++ {
			value += graph[edge(seating[i], seating[i+1])]
		}
		max = math.Max(max, value)
	}

	return max
}
