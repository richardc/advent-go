package day03

import (
	_ "embed"
	"strings"

	"github.com/richardc/advent-go/input"
	"github.com/richardc/advent-go/math"
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
			Day:   3,
			Part1: func(any) any { return solve(puzzle) },
		},
	)
}

type Link struct {
	direction byte
	distance  int
}

func newLink(s string) Link {
	return Link{
		direction: s[0],
		distance:  input.MustAtoi(s[1:]),
	}
}

type Point [2]int

func solve(puzzle string) int {
	wires := slices.Map(input.Lines(puzzle), func(s string) []Link { return slices.Map(strings.Split(s, ","), newLink) })
	points := map[Point]int{}
	for i, wire := range wires {
		bit := i + 1
		x, y := 0, 0
		for _, link := range wire {
			for i := 0; i < link.distance; i++ {
				switch link.direction {
				case 'D':
					y--
				case 'U':
					y++
				case 'L':
					x--
				case 'R':
					x++
				}
				points[Point{x, y}] |= bit
			}
		}
	}

	crossed := slices.Filter(maps.Keys(points), func(p Point) bool { return points[p] == 3 })
	return slices.Min(slices.Map(crossed, func(p Point) int { return math.Abs(p[0]) + math.Abs(p[1]) }))
}
